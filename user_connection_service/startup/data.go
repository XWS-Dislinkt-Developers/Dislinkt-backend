package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
)

var userConnections = []*domain.UserConnection{
	{
		UserId:  1,
		Private: true,
		Connections: []int{
			3,
		},
		Requests: []int{4},
	},
	{
		UserId:  2,
		Private: false,
		Connections: []int{
			3, 4,
		},
		Requests: []int{},
	},
	{
		UserId:  3,
		Private: false,
		Connections: []int{
			2, 1,
		},
		Requests: []int{},
	},
	{
		UserId:  4,
		Private: true,
		Connections: []int{
			2,
		},
		Requests: []int{},
	},
}
