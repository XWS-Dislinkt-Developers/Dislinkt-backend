package api

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	"net/http"
	"strings"
)

type UserHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	auth_service *application.AuthService
	loggerInfo   *logg.Logger
	loggerError  *logg.Logger
}

func NewUserHandler(auth_service *application.AuthService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserHandler {
	return &UserHandler{
		auth_service: auth_service,
		loggerInfo:   loggerInfo,
		loggerError:  loggerError,
	}
}

func (handler *UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user domain.User
	user.Username = request.User.Username
	user.Email = request.User.Email
	user.Password = handler.auth_service.HashPassword(request.User.Password)
	user.Role = "user"
	if len(strings.TrimSpace(user.Username)) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Username can't be empty.",
		}, nil
	}
	if len(strings.TrimSpace(user.Email)) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Email can't be empty.",
		}, nil
	}
	if len(strings.TrimSpace(user.Password)) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Password can't be empty.",
		}, nil
	}
	if request.User.Password != request.User.ConfirmPassword {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Password mismatch",
		}, nil
	}
	_, err := handler.auth_service.GetByUsername(user.Username)
	if err == nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Username is already taken",
		}, nil
	}

	if !handler.auth_service.IsPasswordValid(strings.TrimSpace(user.Password)) {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error: "Your password must be at least 10 characters long, containing at least 1 lowercase and 1 uppercase letter," +
				" at least 1 special character(~`!@#$%^&* etc.) and a number.",
		}, nil
	}

	if !handler.auth_service.CheckForCommonPasswords(strings.TrimSpace(user.Password)) {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Your password is on the list of most commonly used passwords. Please, pick another, more complex password.",
		}, nil
	}

	registratedUser := handler.auth_service.Create(&user)

	handler.auth_service.SendEmailForUserAuthentication(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		UserId: int64(registratedUser.ID),
	}, nil

}

func (handler *UserHandler) ConfirmAccount(ctx context.Context, req *pb.ConfirmAccountRequest) (*pb.ConfirmAccountResponse, error) {

	user, _ := handler.auth_service.ConfirmAccount(req.Token)

	return &pb.ConfirmAccountResponse{
		Email:    user.Email,
		Response: "Posetite nas sajt i ulogujte se: http://localhost:4200",
	}, nil
}

func (handler *UserHandler) PasswordRecoveryRequest(ctx context.Context, req *pb.PasswordRecoveryReq) (*pb.PasswordRecoveryResponse, error) {

	err := handler.auth_service.PasswordRecoveryRequest(req.Email.Email)

	if err != nil {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.PasswordRecoveryResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil

}

func (handler *UserHandler) PasswordRecovery(ctx context.Context, req *pb.ChangePasswordWithCodeRequest) (*pb.PasswordRecoveryResponse, error) {

	if len(strings.TrimSpace(req.ChangePassword.Code)) == 0 {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  "Code can't be empty.",
		}, nil
	}

	if len(strings.TrimSpace(req.ChangePassword.Password)) == 0 {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  "Password can't be empty.",
		}, nil
	}

	if req.ChangePassword.Password != req.ChangePassword.ConfirmPassword {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  "Password mismatch",
		}, nil
	}

	if !handler.auth_service.IsPasswordValid(strings.TrimSpace(req.ChangePassword.Password)) {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error: "Your password must be at least 10 characters long, containing at least 1 lowercase and 1 uppercase letter," +
				" at least 1 special character(~`!@#$%^&* etc.) and a number.",
		}, nil
	}

	if !handler.auth_service.CheckForCommonPasswords(strings.TrimSpace(req.ChangePassword.Password)) {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  "Your password is on the list of most commonly used passwords. Please, pick another, more complex password.",
		}, nil
	}

	err, user := handler.auth_service.PasswordRecovery(req.ChangePassword.Code, req.ChangePassword.Password)

	if err != "" {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  err,
		}, nil
	}

	return &pb.PasswordRecoveryResponse{
		Status: http.StatusOK,
		Error:  "Password successfully changed.",
		Email:  user.Email,
	}, nil
}

func (handler *UserHandler) PasswordlessLoginRequest(ctx context.Context, req *pb.PasswordlessLoginReq) (*pb.PasswordRecoveryResponse, error) {

	user, _ := handler.auth_service.GetByEmail(req.Email.Email)
	err := handler.auth_service.PasswordlessLoginRequest(user)

	if err != nil {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	return &pb.PasswordRecoveryResponse{
		Status: http.StatusOK,
		Error:  "Email sent successfully!",
	}, nil
}

func (handler *UserHandler) PasswordlessLogin(ctx context.Context, req *pb.PasswordlessLogRequest) (*pb.LoginResponse, error) {

	user, res := handler.auth_service.PasswordlessLogin(req.Code.Code)

	if user != nil {
		token, _ := handler.auth_service.GenerateToken(user)

		return &pb.LoginResponse{
			Status:   http.StatusOK,
			Token:    token,
			Error:    res,
			Username: user.Username,
			Id:       int64(user.ID),
			Role:     user.Role,
		}, nil
	}

	return &pb.LoginResponse{
		Status:   http.StatusBadRequest,
		Token:    "",
		Error:    res,
		Username: "",
		Id:       0,
		Role:     "",
	}, nil
}

func (handler *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := handler.auth_service.GetByUsername(req.UserData.Username)
	println("[handler *UserHandler]Login")

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	if user.IsItConfirmed == false {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User's account is not confirmed!",
		}, nil
	}

	match := handler.auth_service.CheckPasswordHash(req.UserData.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Password is wrong!",
		}, nil
	}

	token, _ := handler.auth_service.GenerateToken(user)

	return &pb.LoginResponse{
		Status:   http.StatusOK,
		Token:    token,
		Username: user.Username,
		Id:       int64(user.ID),
		Role:     user.Role,
	}, nil
}

func (handler *UserHandler) FindUser(ctx context.Context, request *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	User, err := handler.auth_service.GetByUsername(request.FindUser.Username)
	if err != nil || User == nil {
		return nil, err
	}
	UserPb := mapUser(User)

	response := &pb.FindUserResponse{
		User: UserPb,
	}

	return response, nil
}

func (handler *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := handler.auth_service.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	user, err := handler.auth_service.Get(claims.Id)
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
