package domain

type UserConnectionStore interface {
	GetByUserId(id int) (*UserConnection, error)
	GetAll() ([]*UserConnection, error)
	Insert(userPost *UserConnection) error
	DeleteAll()
}
