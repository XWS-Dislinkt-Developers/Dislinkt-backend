package api

import (
	"context"
	"encoding/json"
	"fmt"
	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_conn "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	pb_post "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	app_conn "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/application"
	app_post "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type UserPostHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService

	pb_post.UnimplementedUserPostServiceServer
	post_service *app_post.UserPostService

	pb_conn.UnimplementedUserConnectionServiceServer
	conn_service *app_conn.UserConnectionService

	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserPostHandler(post_service *app_post.UserPostService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserPostHandler {

	return &UserPostHandler{
		post_service: post_service,
		loggerInfo:   loggerInfo,
		loggerError:  loggerError,
	}
}

//USER
func (handler *UserPostHandler) CreateUserPost(ctx context.Context, request *pb_post.CreateUserPostRequest) (*pb_post.CreateUserPostResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)
	//var creatorUserId = claims.Id

	userPost := mapNewUserPost(request.UserPost, claims.Id)
	err := handler.post_service.Create(userPost)

	if err != nil {
		handler.loggerInfo.Logger.Errorf(err.Error())
		return nil, err
	}

	handler.loggerInfo.Logger.Infof("User_post_grpc_handler: CreateUserPost - User with id " + strconv.Itoa(claims.Id) + " created new post.")

	return &pb_post.CreateUserPostResponse{
		UserPost: mapUserPost(userPost),
	}, nil
}

func (handler *UserPostHandler) Get(ctx context.Context, request *pb_post.GetRequest) (*pb_post.GetResponse, error) {
	//return nil, status.Error(codes.Unauthenticated, "Your role doesn't allow this method")
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
func (handler *UserPostHandler) GetPostsForFeed(ctx context.Context, request *pb_post.GetAllRequest) (*pb_post.GetAllResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)
	//PROVERA ULOGE
	//if handler.auth_service.CheckIfUser(claims.Role) == false {
	//	return nil, status.Error(codes.Unauthenticated, "Your role doesn't allow you this method.")
	//}
	IdLoggedUser := claims.Id

	AllUserConnections := make([]int, 0)
	feedPosts := make([]*domain.UserPost, 0)
	resp, err := http.Get("http://localhost:8000/userConnections")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
	var responsenew ResponseNew
	err = json.Unmarshal(body, &responsenew)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}

	AllConnections, _ := handler.conn_service.GetAll()

	for _, userConnection := range AllConnections {
		if userConnection.UserId == IdLoggedUser {
			AllUserConnections = userConnection.Connections
		}
	}

	for _, idConnection := range AllUserConnections {
		Posts, _ := handler.post_service.GetUserPosts(idConnection)
		for _, c := range Posts {
			feedPosts = append(feedPosts, c)
		}
	}

	response := &pb_post.GetAllResponse{
		UserPosts: []*pb_post.UserPost{},
	}
	for _, UserPost := range feedPosts {
		current := mapUserPost(UserPost)
		response.UserPosts = append(response.UserPosts, current)
	}
	return response, nil

}
func (handler *UserPostHandler) GetAll(ctx context.Context, request *pb_post.GetAllRequest) (*pb_post.GetAllResponse, error) {
	//header, _ := extractHeader(ctx, "authorization")
	//var prefix = "Bearer "
	//var token = strings.TrimPrefix(header, prefix)
	//claims, _ := handler.auth_service.ValidateToken(token)
	//PROVERA ULOGE
	//if handler.auth_service.CheckIfUser(claims.Role) == false {
	//	return nil, status.Error(codes.Unauthenticated, "Your role doesn't allow you this method.")
	//}

	handler.loggerInfo.Logger.Infof("User_post_grpc_handler: GetAll | id user ---> 3|")

	userPosts, err := handler.post_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf(err.Error())
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

// TODO:AddReactionToUserPost
func (handler *UserPostHandler) AddReactionToUserPost(ctx context.Context, request *pb_post.AddReactionRequest) (*pb_post.GetResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	newReaction := mapNewReactionToUserPost(request, claims.Id)
	postId, _ := primitive.ObjectIDFromHex(request.AddReaction.PostId)

	UserPost, _ := handler.post_service.AddReaction(newReaction, postId)
	handler.loggerInfo.Logger.Infof("User_post_grpc_handler: AddReactionToUserPost - User with id " + strconv.Itoa(claims.Id) + " reacted to post with id " + request.AddReaction.PostId + ".")

	UserPostPb := mapUserPost(UserPost)
	response := &pb_post.GetResponse{
		UserPost: UserPostPb,
	}
	return response, nil

}

//TODO:AddCommentToUserPost()
func (handler *UserPostHandler) AddComment(ctx context.Context, request *pb_post.AddCommentRequest) (*pb_post.GetResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)
	newComment := mapNewCommentToUserPost(request, claims.Id)

	postId, _ := primitive.ObjectIDFromHex(request.AddComment.IdPost)
	UserPost, _ := handler.post_service.AddComment(newComment, postId)
	handler.loggerInfo.Logger.Infof("User_post_grpc_handler: AddCommentToUserPost - User with id " + strconv.Itoa(claims.Id) + " added comment to post with id " + request.AddComment.IdPost + ".")

	UserPostPb := mapUserPost(UserPost)
	response := &pb_post.GetResponse{
		UserPost: UserPostPb,
	}
	return response, nil

}

func (handler *UserPostHandler) GetUserPosts(ctx context.Context, request *pb_post.GetUserPostsRequest) (*pb_post.GetAllResponse, error) {
	id := int(request.Id)
	userPosts, err := handler.post_service.GetUserPosts(id)
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

type ResponseNew struct {
	UserConnections []UserConnection `json:"userConnections"`
}

//json.RawMessage
type UserConnection struct {
	UserId      string            `json:"userId"`
	Private     bool              `json:"private"`
	Connections []json.RawMessage `json:"connections"`
	Requests    []json.RawMessage `json:"requests"`
}

type Connection struct {
	con []string
}
