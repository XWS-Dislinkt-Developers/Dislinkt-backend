package api

import (
	"context"
	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/application"
	app_user "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
	"time"
)

type UsersHandler struct {
	loggerInfo  *logg.Logger
	loggerError *logg.Logger

	pb.UnimplementedUserServiceServer
	service *app_user.UserService

	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService
}

func NewUsersHandler(service *application.UserService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UsersHandler {
	return &UsersHandler{
		service:     service,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (handler *UsersHandler) UpdatePersonalData(ctx context.Context, request *pb.UpdatePersonalDataRequest) (*pb.UpdatePersonalDataResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	var dto domain.UpdateUserDto
	dto.Username = request.UpdateUserData.Username
	dto.Email = request.UpdateUserData.Email
	dto.Gender = request.UpdateUserData.Gender
	dto.Biography = request.UpdateUserData.Biography
	dto.PhoneNumber = request.UpdateUserData.PhoneNumber
	dto.Name = request.UpdateUserData.Name
	myDate, err := time.Parse("2006-01-02", request.UpdateUserData.DateOfBirth)
	if err == nil {
		dto.DateOfBirth = myDate
	}

	_, err = handler.service.UpdateUser(dto, claims.Id)
	if err != nil {
		return &pb.UpdatePersonalDataResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.UpdatePersonalDataResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UsersHandler) UpdateUserWorkEducation(ctx context.Context, request *pb.UpdateUserWAERequest) (*pb.UpdateUserWAEResponse, error) {

	header, err := extractHeader(ctx, "authorization")
	if err != nil {
		return &pb.UpdateUserWAEResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, err := handler.auth_service.ValidateToken(token)
	if err != nil {
		return &pb.UpdateUserWAEResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var dto domain.UpdateUserWAEDto
	dto.Work = request.UpdateUserData.Work
	dto.Education = request.UpdateUserData.Education

	handler.service.UpdateUserWAE(dto, claims.Id)

	return &pb.UpdateUserWAEResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UsersHandler) UpdateUserSkillsInterests(ctx context.Context, request *pb.UpdateUserSAIRequest) (*pb.UpdateUserSAIResponse, error) {

	header, err := extractHeader(ctx, "authorization")
	if err != nil {
		return &pb.UpdateUserSAIResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, err := handler.auth_service.ValidateToken(token)
	if err != nil {
		return &pb.UpdateUserSAIResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var dto domain.UpdateUserSAIDto
	dto.Skills = request.UpdateUserData.Skills
	dto.Interests = request.UpdateUserData.Interests

	handler.service.UpdateUserSAI(dto, claims.Id)

	return &pb.UpdateUserSAIResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UsersHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
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

func (handler *UsersHandler) CreateUser(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user domain.User
	user.Username = request.User.Username
	user.Name = request.User.Name
	user.Email = request.User.Email
	user.Password = handler.auth_service.HashPassword(request.User.Password)
	user.Gender = request.User.Gender
	user.IsPrivateProfile = false
	user.Role = "user"
	handler.service.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UsersHandler) ConfirmAccount(ctx context.Context, request *pb.ConfirmAccountRequest) (*pb.ConfirmAccountResponse, error) {

	handler.service.ConfirmAccount(request.Email)

	return &pb.ConfirmAccountResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UsersHandler) ChangePassword(ctx context.Context, request *pb.ChangePasswordReq) (*pb.ConfirmAccountResponse, error) {

	handler.service.ChangePassword(request.ChangePassword.Email, request.ChangePassword.Password)

	return &pb.ConfirmAccountResponse{
		Status: http.StatusOK,
		Error:  "",
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
