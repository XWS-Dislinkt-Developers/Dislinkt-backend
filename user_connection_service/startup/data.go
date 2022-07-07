package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
)

var userConnections = []*domain.UserConnection{
	{
		UserId:  13,
		Private: true,
		Connections: []int{
			12, 14,
		},
		Requests:        []int{5},
		WaitingResponse: []int{6, 7},
		Blocked:         []int{9},
	},
	{
		UserId:  12,
		Private: false,
		Connections: []int{
			3, 4,
		},
		Requests:        []int{},
		WaitingResponse: []int{},
		Blocked:         []int{},
	},
	{
		UserId:  14,
		Private: false,
		Connections: []int{
			2, 1,
		},
		Requests:        []int{},
		WaitingResponse: []int{},
		Blocked:         []int{},
	},
	{
		UserId:  4,
		Private: true,
		Connections: []int{
			2,
		},
		Requests:        []int{},
		WaitingResponse: []int{},
		Blocked:         []int{},
	},
}
