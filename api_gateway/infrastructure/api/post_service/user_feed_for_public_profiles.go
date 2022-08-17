package post_service

import (
	"context"
	"encoding/json"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pbCPost "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	pbUser "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"sort"
)

type UserFeedForPublicProfilesHandler struct {
	postClientAddress string
	userClientAddress string
}

func NewUserFeedForPublicProfilesHandler(postClientAddress, userClientAddress string) *UserFeedForPublicProfilesHandler {
	return &UserFeedForPublicProfilesHandler{
		postClientAddress: postClientAddress,
		userClientAddress: userClientAddress,
	}
}

func (handler *UserFeedForPublicProfilesHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/publicUserFeed", handler.HandleUserFeed)
	if err != nil {
		panic(err)
	}
}

func (handler *UserFeedForPublicProfilesHandler) HandleUserFeed(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	//TODO: endpoint koji iz servisa users vraca sve public profile
	usersClient := services.NewUserClient(handler.userClientAddress)
	response, err := usersClient.GetAllPublicProfiles(context.TODO(), &pbUser.GetAllRequest{})
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
	for _, c := range response.Users.UserId {
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
