package domain

type ConfirmationTokenStore interface {
	Insert(confirmationToken *ConfirmationToken) error
	GetByUserId(userId int) (*ConfirmationToken, error)
	GetByConfirmationToken(token string) (*ConfirmationToken, error)
	GetAll() (*[]ConfirmationToken, error)
	Delete(userId int)
}
