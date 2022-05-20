package domain

type UserStore interface {
	Insert(user *User) error
	Get(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	UpdateUser(dto UpdateUserDto, userID int) (*User, error)
	UpdateUserWorkAndEducation(dto UpdateUserWAEDto, userId int) (*User, error)
	UpdateUserSkillsAndInterests(dto UpdateUserSAIDto, userId int) (*User, error)
	GetAll() (*[]User, error)
	DeleteAll()
	ConfirmAccount(idUser int) (*User, error)
	GetById(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}
