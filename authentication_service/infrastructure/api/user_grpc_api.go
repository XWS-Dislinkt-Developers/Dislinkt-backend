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
	"time"
)

type UserHandler struct {
	service *application.UserService
	pb.UnimplementedAuthenticationServiceServer
	auth_service *application.AuthService
}

func NewUserHandler(service *application.UserService, auth_service *application.AuthService) *UserHandler {
	return &UserHandler{
		service:      service,
		auth_service: auth_service,
	}
}

func (handler *UserHandler) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var user domain.User

	user.Username = request.User.Username
	user.Name = request.User.Name
	user.Email = request.User.Email
	user.Password = handler.auth_service.HashPassword(request.User.Password)
	user.Gender = request.User.Gender
	user.IsPrivateProfile = false
	if len(strings.TrimSpace(user.Username)) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Username can't be empty.",
		}, nil
	}
	if len(strings.TrimSpace(user.Name)) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Name can't be empty.",
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
	if len(strings.TrimSpace(user.Gender)) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Gender can't be empty.",
		}, nil
	}
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

	//ovde pozvati fiju za proveru da li je slaba lozinka

	handler.service.Create(&user)

	handler.auth_service.SendEmailForUserAuthentication(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil

}

func (handler *UserHandler) ConfirmAccount(ctx context.Context, req *pb.ConfirmAccountRequest) (*pb.ConfirmAccountResponse, error) {

	handler.auth_service.ConfirmAccount(req.Token)

	return &pb.ConfirmAccountResponse{
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

	//ovde pozvati fiju za proveru da li je slaba lozinka

	err := handler.auth_service.PasswordRecovery(req.ChangePassword.Code, req.ChangePassword.Password)

	if err != "" {
		return &pb.PasswordRecoveryResponse{
			Status: http.StatusBadRequest,
			Error:  err,
		}, nil
	}

	return &pb.PasswordRecoveryResponse{
		Status: http.StatusOK,
		Error:  "Password successfully changed.",
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

//------------------------------------------------------------USER SERVICE------------------------------------------------------------

func (handler *UserHandler) UpdatePersonalData(ctx context.Context, request *pb.UpdatePersonalDataRequest) (*pb.UpdatePersonalDataResponse, error) {

	print("ucitava se sendler")
	header, err := extractHeader(ctx, "authorization")
	if err != nil {
		return &pb.UpdatePersonalDataResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, err := handler.auth_service.ValidateToken(token)
	if err != nil {
		return &pb.UpdatePersonalDataResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	println("id je :", claims.Id)

	var dto domain.UpdateUserDto
	dto.Username = request.UpdateUserData.Username
	dto.Email = request.UpdateUserData.Email
	dto.Gender = request.UpdateUserData.Gender
	dto.Biography = request.UpdateUserData.Biography
	dto.PhoneNumber = request.UpdateUserData.Biography
	dto.Name = request.UpdateUserData.Name
	//dto.DateOfBirth = request.UpdateUserData.DateOfBirth
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

func (handler *UserHandler) UpdateUserWorkEducation(ctx context.Context, request *pb.UpdateUserWAERequest) (*pb.UpdateUserWAEResponse, error) {

	print("ucitava se sendler")
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
	println("id je :", claims.Id)

	var dto domain.UpdateUserWAEDto
	dto.Work = request.UpdateUserData.Work
	dto.Education = request.UpdateUserData.Education

	handler.service.UpdateUserWAE(dto, claims.Id)

	return &pb.UpdateUserWAEResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UserHandler) UpdateUserSkillsInterests(ctx context.Context, request *pb.UpdateUserSAIRequest) (*pb.UpdateUserSAIResponse, error) {

	print("ucitava se sendler")
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
	println("id je :", claims.Id)

	var dto domain.UpdateUserSAIDto
	dto.Skills = request.UpdateUserData.Skills
	dto.Interests = request.UpdateUserData.Interests

	handler.service.UpdateUserSAI(dto, claims.Id)

	return &pb.UpdateUserSAIResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UserHandler) FindUser(ctx context.Context, request *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	print("USAO SAM U HENDLER ")
	print(request.FindUser.Username)
	User, err := handler.service.GetByUsername(request.FindUser.Username)
	if err != nil || User == nil {
		return nil, err
	}
	UserPb := mapUser(User)

	response := &pb.FindUserResponse{
		User: UserPb,
	}

	return response, nil
}

func (handler *UserHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {

	//header, _ := extractHeader(ctx, "authorization")
	//var prefix = "Bearer "
	//var token = strings.TrimPrefix(header, prefix)
	//claims, _ := handler.auth_service.ValidateToken(token)
	//println("id je :", claims.Id)

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
