package domain

type UserStore interface {
	Insert(user *User) error
	GetAll() (*[]User, error)
	DeleteAll()
}
