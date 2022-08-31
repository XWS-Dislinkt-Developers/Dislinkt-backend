package domain

type GraphConnectionStore2 interface {
	GetByUserId(id int) (*UserConnection, error)
	GetAll() ([]*UserConnection, error)
	Insert(userConnection *UserConn) error
	DeleteAll()
	UpdateRequestConnection(userConnection *UserConnection)
	UpdateConnections(userConnection *UserConnection, loggedUserConnection *UserConnection)
	UpdateBlockedConnection(connection *UserConnection)
	UpdateWaitingResponseConnection(connection *UserConnection)
}
