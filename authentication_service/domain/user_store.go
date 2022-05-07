package domain

type UserStore interface {
	Insert(user *User) error
	Get(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	UpdateUser(dto UpdateUserDto) (*User, error)
	UpdateUserWorkAndEducation(dto UpdateUserWAEDto, userId int) (*User, error)
	UpdateUserSkillsAndInterests(dto UpdateUserSAIDto, userId int) (*User, error)
	GetAll() (*[]User, error)
	DeleteAll()
}
