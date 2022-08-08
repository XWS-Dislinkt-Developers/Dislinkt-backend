package authentication_service

import (
	"context"
	"encoding/json"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/domain"
	services "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	userConnectionPb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	userPb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"io"
	"net/http"
)

type RegisterUserHandler struct {
	authClientAddress           string
	userClientAddress           string
	userConnectionClientAddress string
}

func NewRegisterUserHandler(authClientAddress, userClientAddress, userConnectionClientAddress string) *RegisterUserHandler {
	return &RegisterUserHandler{
		authClientAddress:           authClientAddress,
		userClientAddress:           userClientAddress,
		userConnectionClientAddress: userConnectionClientAddress,
	}
}

func (handler *RegisterUserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/registerUser", handler.HandleRegisterUser)
	if err != nil {
		panic(err)
	}
}

func (handler *RegisterUserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	registerRequestJson, err := decodeBodyToRegisterRequest(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Unable to decode request body!"))
		return
	}

	registerUserRequestPb := mapRegisterUserRequestPb(registerRequestJson)
	authClient := services.NewAuthenticationClient(handler.authClientAddress)
	response, errAuth := authClient.Register(context.TODO(), registerUserRequestPb)
	ret, _ := json.Marshal(response)
	if errAuth != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ret)
		return
	}

	if response.Error == "" {
		registerUserRequestUserPb := mapRegisterUserRequestUserPb(registerRequestJson, response.UserId)
		userClient := services.NewUserClient(handler.userClientAddress)
		_, errUser := userClient.CreateUser(context.TODO(), registerUserRequestUserPb)
		if errUser != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		registerUserConnectionRequestPb := mapRegisterUserConnectionRequestPb(registerRequestJson, response.UserId)
		userConnectionClient := services.NewConnectionClient(handler.userConnectionClientAddress)
		_, errUserConnection := userConnectionClient.RegisterUserConnection(context.TODO(), registerUserConnectionRequestPb)
		if errUserConnection != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

	}

	w.WriteHeader(http.StatusOK)
	w.Write(ret)
	return
}

// Authentication service mapper
func mapRegisterUserRequestPb(user *domain.User) *pb.RegisterRequest {
	var registerUserRequestPb = &pb.RegisterRequest{
		User: &pb.UserRegistrationData{
			Name:            user.Name,
			Username:        user.Username,
			Password:        user.Password,
			ConfirmPassword: user.ConfirmPassword,
			Email:           user.Email,
			Gender:          user.Gender,
			DateOfBirth:     user.DateOfBirth,
		},
	}
	return registerUserRequestPb
}

// User service mapper
func mapRegisterUserRequestUserPb(user *domain.User, userId int64) *userPb.RegisterRequest {
	registerUserRequestPb := &userPb.RegisterRequest{
		User: &userPb.User{
			UserId:      userId,
			Name:        user.Name,
			Username:    user.Username,
			Password:    user.Password,
			Email:       user.Email,
			Gender:      user.Gender,
			DateOfBirth: user.DateOfBirth,
		},
	}
	return registerUserRequestPb
}

// UserConnection service mapper
func mapRegisterUserConnectionRequestPb(user *domain.User, userId int64) *userConnectionPb.RegisterRequest {
	registerUserRequestPb := &userConnectionPb.RegisterRequest{
		IdUser:      userId,
		IsItPrivate: false,
	}
	return registerUserRequestPb
}

func decodeBodyToRegisterRequest(r io.Reader) (*domain.User, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var registerRequest domain.User
	if err := dec.Decode(&registerRequest); err != nil {
		return nil, err
	}
	return &registerRequest, nil
}
