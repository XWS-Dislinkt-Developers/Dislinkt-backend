package post_service

import (
	"context"
	"encoding/json"
	authService "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/authentication_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pbConn "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	pbCPost "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"sort"
	"strings"
)

type UserFeedHandler struct {
	postClientAddress       string
	connectionClientAddress string
}

func NewUserFeedHandler(postClientAddress, connectionClientAddress string) *UserFeedHandler {
	return &UserFeedHandler{
		postClientAddress:       postClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserFeedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/userFeed", handler.HandleUserFeed)
	if err != nil {
		panic(err)
	}
}
func (handler *UserFeedHandler) HandleUserFeed(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	header := r.Header.Get("Authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := authService.ValidateToken(token)

	//TODO: endpoint koji iz servisa konekcija vraca sve konekcije za ulogovanog korisnika
	connectionClient := services.NewConnectionClient(handler.connectionClientAddress)
	response, err := connectionClient.GetConnectionsByUser(context.TODO(), &pbConn.UserIdRequest{
		IdUser: int64(claims.Id),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//TODO: endpoint koji vraca sve postove iz servisa postova
	postClient := services.NewPostClient(handler.postClientAddress)
	posts, errPost := postClient.GetAll(context.TODO(), &pbCPost.GetAllRequest{})
	if errPost != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//TODO: uzeti sve postove za svaku konekciju, sortirati ih po datumu i vratiti na front
	feed := make([]domain.UserPost, 0)
	response.Connections.Connections = append(response.Connections.Connections, int64(claims.Id)) // The user can see his own posts
	for _, c := range response.Connections.Connections {
		for _, p := range posts.UserPosts {
			if c == p.UserId {
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
		}
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
func converterInt64ToIntArray(likes []int64) []int {
	ret := make([]int, 0)
	for i := range likes {
		ret = append(ret, int(likes[i]))
	}
	return ret
}

func getComments(comments []*pbCPost.Comment) []domain.Comment {
	ret := make([]domain.Comment, 0)
	for _, c := range comments {
		comment := domain.Comment{
			UserId:    int(c.UserId),
			Text:      c.Text,
			CreatedAt: c.CreatedAt.AsTime().Local(),
		}
		ret = append(ret, comment)
	}
	return ret
}
