package post_service

import (
	"context"
	"encoding/json"
	authService "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/authentication_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pbConn "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	pbPost "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	pbUser "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type UserPostsProfileHandler struct {
	userClientAddress       string
	postClientAddress       string
	connectionClientAddress string
}

func NewUserPostsProfileHandler(userClientAddress, postClientAddress, connectionClientAddress string) *UserPostsProfileHandler {
	return &UserPostsProfileHandler{
		userClientAddress:       userClientAddress,
		postClientAddress:       postClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserPostsProfileHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/userPostsProfile/{idUser}", handler.HandleUserPosts)
	if err != nil {
		panic(err)
	}
}

func (handler *UserPostsProfileHandler) HandleUserPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	idUser, _ := strconv.ParseInt(pathParams["idUser"], 0, 10)
	idLoggedUser := getLoggedUser(r)

	userClient := services.NewUserClient(handler.userClientAddress)
	response, _ := userClient.IsProfilePrivate(context.TODO(), &pbUser.UserIdRequest{
		IdUser: idUser,
	})

	postClient := services.NewPostClient(handler.postClientAddress)
	posts, errPost := postClient.GetUserPosts(context.TODO(), &pbPost.GetUserPostsRequest{
		Id: idUser,
	})
	if errPost != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errPost.Error()))
		return
	}

	isProfilePrivate := response.IsProfilePrivate
	if idLoggedUser == -1 {
		if isProfilePrivate {
			returnEmptyListOfPosts(w)
			return
		} else {
			returnPosts(w, posts)
			return
		}
	}

	connectionClient := services.NewConnectionClient(handler.connectionClientAddress)
	connections, err := connectionClient.GetConnectionsByUser(context.TODO(), &pbConn.UserIdRequest{
		IdUser: int64(idLoggedUser),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if checkConnection(connections, int(idUser)) {
		returnPosts(w, posts)
		return
	} else {
		returnEmptyListOfPosts(w)
		return
	}

}

func returnPosts(w http.ResponseWriter, posts *pbPost.GetAllResponse) {
	feed := make([]domain.UserPost, 0)
	for _, p := range posts.UserPosts {
		po := domain.UserPost{
			Id:        p.Id,
			UserId:    int(p.UserId),
			Text:      p.Text,
			ImagePath: p.ImagePath,
			CreatedAt: p.CreatedAt.AsTime().Local(),
			Comments:  getComments(p.Comments),
			Likes:     converterInt64ToIntArray(p.Likes),
			Dislikes:  converterInt64ToIntArray(p.Dislikes),
		}
		feed = append(feed, po)
	}
	sort.Slice(feed, func(i, j int) bool {
		return feed[i].CreatedAt.After(feed[j].CreatedAt)
	})

	ret := &domain.Fs{
		Feed: feed,
	}
	rt, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(rt)
}

func returnEmptyListOfPosts(w http.ResponseWriter) {
	feed := make([]domain.UserPost, 0)
	ret := &domain.Fs{
		Feed: feed,
	}
	rt, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(rt)
}

func checkConnection(connections *pbConn.Connections, userId int) bool {
	for _, c := range connections.Connections.Connections {
		if int(c) == userId {
			return true
		}
	}
	return false
}

func getLoggedUser(r *http.Request) int {

	header := r.Header.Get("Authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := authService.ValidateToken(token)

	if claims != nil {
		return claims.Id
	}
	return -1
}
