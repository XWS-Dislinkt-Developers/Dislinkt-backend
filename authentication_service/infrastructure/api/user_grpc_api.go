package api

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type UserHandler struct {
	service *application.UserService
	pb.UnimplementedAuthenticationServiceServer
	auth_service *application.AuthService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)
	println("id je :", claims.Id)

	users, err := handler.service.GetAll()
	if err != nil || *users == nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range *users {
		current := mapUser(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (handler *UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user domain.User
	user.Username = request.User.Username
	user.Name = request.User.Name
	user.Email = request.User.Email
	user.Password = request.User.Password
	user.Gender = request.User.Gender
	if request.User.Password != request.User.ConfirmPassword {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Password mismatch",
		}, nil
	}

	println("[UserHandler]Register:Username " + request.User.Username)
	_, err := handler.service.GetByUsername(user.Username)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Username is already taken",
		}, nil
	}
	handler.service.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil

}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := handler.service.GetByUsername(req.UserData.Username)
	println("[handler *UserHandler]Login")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := req.UserData.Password == user.Password

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := handler.auth_service.GenerateToken(user)

	return &pb.LoginResponse{
		Status:   http.StatusOK,
		Token:    token,
		Username: user.Username,
	}, nil
}

func (handler *UserHandler) UpdateUser(ctx context.Context, request *pb.UpdatePersonalDataRequest) (*pb.UpdatePersonalDataResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)
	println("id je :", claims.Id)

	return nil, nil
}

func (handler *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := handler.auth_service.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user, err := handler.service.Get(claims.Id) //TODO GET USER
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(user.ID),
	}, nil
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
