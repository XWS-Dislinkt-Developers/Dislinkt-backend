package api

import (
	"context"
	"net/http"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
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
	user.Password = request.User.Password

	println("[UserHandler]Register:Username " + request.User.Username)
	handler.service.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil

}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := handler.service.GetByUsername(req.UserData.Username)
	println("Here111111111111111111")
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
		UserId: user.ID,
	}, nil
}
