package api

import (
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
)

func mapUserConnection(userConnection *domain.UserConnection) *pb_connection.UserConnection {
	userConnectionPb := &pb_connection.UserConnection{
		UserId:      int64(userConnection.UserId),
		Private:     userConnection.Private,
		Connections: make([]int64, 0),
		Requests:    make([]int64, 0),
	}
	for _, connection := range userConnection.Connections {
		userConnectionPb.Connections = append(userConnectionPb.Connections, int64(connection))
	}
	for _, request := range userConnection.Requests {
		userConnectionPb.Requests = append(userConnectionPb.Requests, int64(request))
	}
	return userConnectionPb
}
