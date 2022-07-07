package connection_service

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
	"strings"
)

type UserWaitingResponseHandler struct {
	userClientAddress       string
	connectionClientAddress string
}

func NewUserWaitingResponseHandler(userClientAddress, connectionClientAddress string) *UserWaitingResponseHandler {
	return &UserWaitingResponseHandler{
		userClientAddress:       userClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserWaitingResponseHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/waitingResponsesForLoggedUser", handler.HandleUserResponses)
	if err != nil {
		panic(err)
	}
}

func (handler *UserWaitingResponseHandler) HandleUserResponses(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
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
	for _, c := range response.WaitingResponse {
		for _, u := range users.Users {
			if c == u.UserId {
				user := domain.UserDTO{
					ID:       int(u.Id),
					UserId:   int(u.UserId),
					Name:     u.Name,
					Username: u.Username,
				}
				userConnections = append(userConnections, user)
			}
		}
	}

	ret := &domain.Users{
		Users: userConnections,
	}
	rt, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(rt)
}
