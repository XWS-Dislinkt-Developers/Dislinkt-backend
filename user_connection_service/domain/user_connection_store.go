package domain

type UserConnectionStore interface {
	GetByUserId(id int) (*UserConnection, error)
	GetAll() ([]*UserConnection, error)
	Insert(userConnection *UserConnection) error
	DeleteAll()
	AddRequestConnection(userConnection *UserConnection)
	AddConnections(userConnection *UserConnection, loggedUserConnection *UserConnection)
}
