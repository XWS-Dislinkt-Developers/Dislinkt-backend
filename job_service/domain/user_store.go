package domain

type UserDataStore interface {
	GetByUserId(id int) (*UserData, error)
	GetByUserToken(token string) (*UserData, error)
	GetAll() ([]*UserData, error)
	Insert(userData *UserData) error
	AddToken(userData *UserData, newToken string)
}

type JobOfferStore interface {
	GetByUserId(id int) (*JobOffer, error)
	GetAll() ([]*JobOffer, error)
	GetByCompany(company string) ([]*JobOffer, error)
	Insert(jobData *JobOffer) error
}
