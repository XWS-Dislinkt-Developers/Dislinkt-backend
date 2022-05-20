package domain

type PasswordlessLoginStore interface {
	Insert(passwordLogin *PasswordlessLogin) error
	Delete(userId int)
	GetByCode(code string) (passwordlessLogin *PasswordlessLogin, err error)
	GetByUserId(userId int) (*PasswordlessLogin, error)
	GetAll() (*[]PasswordlessLogin, error)
}
