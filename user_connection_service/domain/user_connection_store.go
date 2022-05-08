package domain

type UserConnectionStore interface {
	GetByUserId(id int) (*UserConnection, error)
	GetAll() ([]*UserConnection, error)
	Insert(userConnection *UserConnection) error
	DeleteAll()
	UpdateRequestConnection(userConnection *UserConnection)
	AddConnections(userConnection *UserConnection, loggedUserConnection *UserConnection)
}
