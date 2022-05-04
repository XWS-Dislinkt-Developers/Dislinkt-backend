package api

import (
	"context"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPostHandler struct {
	pb.UnimplementedUserPostServiceServer
	service *application.UserPostService
}

func (handler *UserPostHandler) CreateUserPost(ctx context.Context, request *pb.CreateUserPostRequest) (*pb.CreateUserPostResponse, error) {
	userPost := mapNewUserPost(request.UserPost)
	err := handler.service.Create(userPost)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserPostResponse{
		UserPost: mapUserPost(userPost),
	}, nil
}

func NewUserPostHandler(service *application.UserPostService) *UserPostHandler {
	return &UserPostHandler{
		service: service,
	}
}

func (handler *UserPostHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	UserPost, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	UserPostPb := mapUserPost(UserPost)
	response := &pb.GetResponse{
		UserPost: UserPostPb,
	}
	return response, nil
}

func (handler *UserPostHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	userPosts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		UserPost: []*pb.UserPost{},
	}
	for _, UserPost := range userPosts {
		current := mapUserPost(UserPost)
		response.UserPosts = append(response.UserPosts, current)
	}
	return response, nil
}
