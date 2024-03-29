package domain

type UserConnectionStore interface {
	GetByUserId(id int) (*UserConnection, error)
	GetAll() ([]*UserConnection, error)
	Insert(userConnection *UserConnection) error
	DeleteAll()
	UpdateRequestConnection(userConnection *UserConnection)
	UpdateConnections(userConnection *UserConnection, loggedUserConnection *UserConnection)
	UpdateBlockedConnection(connection *UserConnection)
	UpdateWaitingResponseConnection(connection *UserConnection)

	GetAllUserNotificationsByUserId(id int) ([]*Notification, error)
	InsertNotification(notification *Notification) error
}
