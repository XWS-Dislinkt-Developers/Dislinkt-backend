package api

import (
	"context"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"

	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	app_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/application"
)

type UserConnectionHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService

	pb_connection.UnimplementedUserConnectionServiceServer
	connection_service *app_connection.UserConnectionService
}

func NewUserConnectionHandler(connection_service *app_connection.UserConnectionService) *UserConnectionHandler {
	return &UserConnectionHandler{
		connection_service: connection_service,
	}
}

func (handler *UserConnectionHandler) GetAll(ctx context.Context, request *pb_connection.GetAllRequest) (*pb_connection.GetAllResponse, error) {
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb_connection.GetAllResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
