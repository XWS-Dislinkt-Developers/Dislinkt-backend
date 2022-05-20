package domain

type PasswordRecoveryStore interface {
	Insert(passwordRecovery *PasswordRecovery) error
	Delete(userId int)
	GetByRecoveryCode(recoveryCode string) (passwordRecovery *PasswordRecovery, err error)
	GetByUserId(userId int) (*PasswordRecovery, error)
	GetAll() (*[]PasswordRecovery, error)
}
