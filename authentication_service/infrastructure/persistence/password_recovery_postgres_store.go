package persistence

import (
	"errors"
	ftm "fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"gorm.io/gorm"
)

type PasswordRecoveryPostgresStore struct {
	db *gorm.DB
}

func NewPasswordRecoveryPostgresStore(db *gorm.DB) (domain.PasswordRecoveryStore, error) {
	err := db.AutoMigrate(&domain.PasswordRecovery{})
	if err != nil {
		return nil, err
	}
	return &PasswordRecoveryPostgresStore{db: db}, nil
}

func (store *PasswordRecoveryPostgresStore) GetByRecoveryCode(recoveryCode string) (passwordRecovery *domain.PasswordRecovery, err error) {
	var password *domain.PasswordRecovery
	passwords, err := store.GetAll()
	if err != nil {
		return nil, errors.New("[PasswordRecoveryPostgresStore-GetByRecoveryCode(confToken)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.RecoveryCode == recoveryCode {
			return &p, nil
		}
	}
	if password == nil {
		return nil, errors.New("ERR - [ConfirmationTokenPostgresStore-GetByConfirmationToken(confToken)]: Can't find user with this token: " + recoveryCode)
	}
	return password, nil
}

func (store *PasswordRecoveryPostgresStore) GetByUserId(id int) (*domain.PasswordRecovery, error) {
	var password *domain.PasswordRecovery
	passwords, err := store.GetAll()
	if err != nil {
		return nil, errors.New("[PasswordRecoveryPostgresStore-GetByUserId(id)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.UserId == id {
			return &p, nil
		}
	}
	if password == nil {
		return nil, errors.New("ERR - [PasswordRecoveryPostgresStore-GetByUserId(id)]: Can't find user with this id: " + string(id))
	}
	return password, nil
}

func (store *PasswordRecoveryPostgresStore) Insert(passwordRecovery *domain.PasswordRecovery) error {
	result := store.db.Create(passwordRecovery)
	if result.Error != nil {
		ftm.Println("[PasswordRecoveryPostgresStore-Insert(passwordRecovery)]: Can't insert passwordRecovery.")
		return errors.New("ERR - [PasswordRecoveryPostgresStore-Insert(passwordRecovery)]: Can't insert passwordRecovery. ")
	}
	return nil
}

func (store *PasswordRecoveryPostgresStore) GetAll() (*[]domain.PasswordRecovery, error) {
	var codes []domain.PasswordRecovery
	result := store.db.Find(&codes)
	if result.Error != nil {
		return nil, result.Error
	}
	return &codes, nil
}

func (store *PasswordRecoveryPostgresStore) Delete(idUser int) {
	token, _ := store.GetByUserId(idUser)
	store.db.Delete(&token)
}
