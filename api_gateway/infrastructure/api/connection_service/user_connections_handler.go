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

type UserConnectionsHandler struct {
	userClientAddress       string
	connectionClientAddress string
}

func NewUserConnectionsHandler(userClientAddress, connectionClientAddress string) *UserConnectionsHandler {
	return &UserConnectionsHandler{
		userClientAddress:       userClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserConnectionsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/connectionsForLoggedUser", handler.HandleUserConnections)
	if err != nil {
		panic(err)
	}
}

func (handler *UserConnectionsHandler) HandleUserConnections(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	header := r.Header.Get("Authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := authService.ValidateToken(token)

	connectinClient := services.NewConnectionClient(handler.connectionClientAddress)

	response, err := connectinClient.GetConnectionsByUser(context.TODO(), &pbConn.UserIdRequest{
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
	for _, c := range response.Connections.Connections {
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
