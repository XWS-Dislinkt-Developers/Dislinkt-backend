package persistence

import (
	"errors"
	ftm "fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"gorm.io/gorm"
)

type PasswordlessLoginPostgresStore struct {
	db *gorm.DB
}

func NewPasswordlessLoginPostgresStore(db *gorm.DB) (domain.PasswordlessLoginStore, error) {
	err := db.AutoMigrate(&domain.PasswordlessLogin{})
	if err != nil {
		return nil, err
	}
	return &PasswordlessLoginPostgresStore{db: db}, nil
}

func (store *PasswordlessLoginPostgresStore) GetByCode(code string) (passwordlessLogin *domain.PasswordlessLogin, err error) {
	var password *domain.PasswordlessLogin
	passwords, err := store.GetAll()
	if err != nil {
		return nil, errors.New("[PasswordlessLoginPostgresStore-GetByCode(recoveryCode)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.Code == code {
			return &p, nil
		}
	}
	if password == nil {
		return nil, errors.New("ERR - [PasswordlessLoginPostgresStore-GetByCode(recoveryCode)]: Can't find user with this token: " + code)
	}
	return password, nil
}

func (store *PasswordlessLoginPostgresStore) GetByUserId(id int) (*domain.PasswordlessLogin, error) {
	var password *domain.PasswordlessLogin
	passwords, err := store.GetAll()
	if err != nil {
		return nil, errors.New("[PasswordlessLoginPostgresStore-GetByUserId(id)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.UserId == id {
			return &p, nil
		}
	}
	if password == nil {
		return nil, errors.New("ERR - [PasswordlessLoginPostgresStore-GetByUserId(id)]: Can't find user with this id: " + string(id))
	}
	return password, nil
}

func (store *PasswordlessLoginPostgresStore) Insert(passwordlessLogin *domain.PasswordlessLogin) error {
	result := store.db.Create(passwordlessLogin)
	if result.Error != nil {
		ftm.Println("[PasswordlessLoginPostgresStore-Insert(passwordlessLogin)]: Can't insert passwordlessLogin.")
		return errors.New("ERR - [PasswordlessLoginPostgresStore-Insert(passwordlessLogin)]: Can't insert passwordlessLogin. ")
	}
	return nil
}

func (store *PasswordlessLoginPostgresStore) GetAll() (*[]domain.PasswordlessLogin, error) {
	var codes []domain.PasswordlessLogin
	result := store.db.Find(&codes)
	if result.Error != nil {
		return nil, result.Error
	}
	return &codes, nil
}

func (store *PasswordlessLoginPostgresStore) Delete(idUser int) {
	passwordlessLogin, _ := store.GetByUserId(idUser)
	store.db.Where("user_id = ?", idUser).Delete(&passwordlessLogin)
}
