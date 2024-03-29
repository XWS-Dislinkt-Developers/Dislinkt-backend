package domain

type UserStore interface {
	Insert(user *User) (error, *User)
	Get(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() (*[]User, error)
	DeleteAll()
	ConfirmAccount(idUser int) (*User, error)
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	UpdatePassword(userId int, password string)
	DeleteUser(id int) error
}
