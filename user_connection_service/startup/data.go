package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
)

var userConnections = []*domain.UserConnection{
	{
		UserId: 1,
		Connections: []int{
			3, 1, 5,
		},
	},
}
