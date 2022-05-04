package domain

type UserStore interface {
	Insert(user *User) error
	Get(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() (*[]User, error)
	DeleteAll()
}
