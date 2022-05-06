package api

import (
	"context"

	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_post "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"

	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	app_post "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type UserPostHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService

	pb_post.UnimplementedUserPostServiceServer
	post_service *app_post.UserPostService
}

func NewUserPostHandler(post_service *app_post.UserPostService) *UserPostHandler {
	return &UserPostHandler{
		post_service: post_service,
	}
}

func (handler *UserPostHandler) CreateUserPost(ctx context.Context, request *pb_post.CreateUserPostRequest) (*pb_post.CreateUserPostResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)
	println("id je :", claims.Id)
	//var creatorUserId = claims.Id

	userPost := mapNewUserPost(request.UserPost, claims.Id)
	err := handler.post_service.Create(userPost)
	if err != nil {
		return nil, err
	}
	return &pb_post.CreateUserPostResponse{
		UserPost: mapUserPost(userPost),
	}, nil
}

func (handler *UserPostHandler) Get(ctx context.Context, request *pb_post.GetRequest) (*pb_post.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	UserPost, err := handler.post_service.Get(objectId)
	if err != nil {
		return nil, err
	}
	UserPostPb := mapUserPost(UserPost)
	response := &pb_post.GetResponse{
		UserPost: UserPostPb,
	}
	return response, nil
}

func (handler *UserPostHandler) GetAll(ctx context.Context, request *pb_post.GetAllRequest) (*pb_post.GetAllResponse, error) {
	userPosts, err := handler.post_service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb_post.GetAllResponse{
		UserPosts: []*pb_post.UserPost{},
	}
	for _, UserPost := range userPosts {
		current := mapUserPost(UserPost)
		response.UserPosts = append(response.UserPosts, current)
	}
	return response, nil
}

// TODO: AddReactionToUserPost(), AddCommentToUserPost()

func extractHeader(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no headers in request")
	}

	authHeaders, ok := md[header]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no header in request")
	}

	if len(authHeaders) != 1 {
		return "", status.Error(codes.Unauthenticated, "more than 1 header in request")
	}

	return authHeaders[0], nil
}
