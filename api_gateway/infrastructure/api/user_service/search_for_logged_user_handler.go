package user_service

import (
	"context"
	"encoding/json"
	authService "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/authentication_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pbConn "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	pbUser "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"strconv"
	"strings"
)

type UserSearchForLoggedUserHandler struct {
	userClientAddress       string
	connectionClientAddress string
}

func NewUserSearchForLoggedUserHandler(userClientAddress, connectionClientAddress string) *UserSearchForLoggedUserHandler {
	return &UserSearchForLoggedUserHandler{
		userClientAddress:       userClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserSearchForLoggedUserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/searchUsersForLoggedUser/{name}", handler.HandleUserSearch)
	if err != nil {
		panic(err)
	}
}

func (handler *UserSearchForLoggedUserHandler) HandleUserSearch(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	name := pathParams["name"]

	header := r.Header.Get("Authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := authService.ValidateToken(token)

	connectionClient := services.NewConnectionClient(handler.connectionClientAddress)

	response, err := connectionClient.GetById(context.TODO(), &pbConn.UserIdRequest{
		IdUser: int64(claims.Id),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userClient := services.NewUserClient(handler.userClientAddress)
	users, errUsers := userClient.GetAll(context.TODO(), &pbUser.GetAllRequest{})
	if errUsers != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userConnections := make([]domain.UserDTO, 0)
	for _, u := range users.Users {
		isFound := strings.Contains(strings.ToLower(u.Name), strings.ToLower(name))
		if userIsNotBlocked(u.UserId, response.Blocked) && isFound {
			private, _ := strconv.ParseBool(u.IsPrivateProfile)
			user := domain.UserDTO{
				ID:               int(u.Id),
				UserId:           int(u.UserId),
				Name:             u.Name,
				Username:         u.Username,
				Gender:           u.Gender,
				IsPrivateProfile: private,
				Biography:        u.Biography,
			}
			userConnections = append(userConnections, user)
		}
	}

	ret := &domain.Users{
		Users: userConnections,
	}
	rt, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(rt)
}

func userIsNotBlocked(idUser int64, blockedUsers []int64) bool {
	for _, c := range blockedUsers {
		if c == idUser {
			return false
		}
	}
	return true
}
