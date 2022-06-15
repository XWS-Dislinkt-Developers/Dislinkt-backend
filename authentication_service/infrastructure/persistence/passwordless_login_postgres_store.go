package persistence

import (
	"errors"
	ftm "fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	"gorm.io/gorm"
	"strconv"
)

type PasswordlessLoginPostgresStore struct {
	db          *gorm.DB
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewPasswordlessLoginPostgresStore(db *gorm.DB, loggerInfo *logg.Logger, loggerError *logg.Logger) (domain.PasswordlessLoginStore, error) {
	err := db.AutoMigrate(&domain.PasswordlessLogin{})
	if err != nil {
		loggerError.Logger.Errorf("Passwordless_login_postgres_store: NewPasswordlessLoginPostgresStore - failed method ")

		return nil, err
	}
	return &PasswordlessLoginPostgresStore{db: db,
		loggerInfo:  loggerInfo,
		loggerError: loggerError}, nil
}

func (store *PasswordlessLoginPostgresStore) GetByCode(code string) (passwordlessLogin *domain.PasswordlessLogin, err error) {
	var password *domain.PasswordlessLogin
	passwords, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Passwordless_login_postgres_store: GetByCode - failed method - there is no user with this code")
		return nil, errors.New("[PasswordlessLoginPostgresStore-GetByCode(recoveryCode)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.Code == code {
			return &p, nil
		}
	}
	if password == nil {
		store.loggerError.Logger.Errorf("Passwordless_login_postgres_store: GetByCode - failed method - can't find user with this token ")

		return nil, errors.New("ERR - [PasswordlessLoginPostgresStore-GetByCode(recoveryCode)]: Can't find user with this token")
	}
	store.loggerInfo.Logger.Infof("Passwordless_login_postgres_store: GetByCode - user is found")
	return password, nil
}

func (store *PasswordlessLoginPostgresStore) GetByUserId(id int) (*domain.PasswordlessLogin, error) {
	var password *domain.PasswordlessLogin
	passwords, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Passwordless_login_postgres_store: GetByUserId - failed method - can't find user with user id " + strconv.Itoa(id))

		return nil, errors.New("[PasswordlessLoginPostgresStore-GetByUserId(id)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.UserId == id {
			return &p, nil
		}
	}
	if password == nil {
		store.loggerError.Logger.Errorf("Passwordless_login_postgres_store: GetByUserId - failed method - can't find user with user id " + strconv.Itoa(id))

		return nil, errors.New("ERR - [PasswordlessLoginPostgresStore-GetByUserId(id)]: Can't find user with this id: " + string(id))
	}
	store.loggerInfo.Logger.Infof("Passwordless_login_postgres_store: GetByUserId - user with id  " + strconv.Itoa(id) + " is found")

	return password, nil
}

func (store *PasswordlessLoginPostgresStore) Insert(passwordlessLogin *domain.PasswordlessLogin) error {
	result := store.db.Create(passwordlessLogin)
	if result.Error != nil {
		store.loggerError.Logger.Errorf("Passwordless_login_postgres_store: Insert - failed method - can't save passwordles login ")

		ftm.Println("[PasswordlessLoginPostgresStore-Insert(passwordlessLogin)]: Can't insert passwordlessLogin.")
		return errors.New("ERR - [PasswordlessLoginPostgresStore-Insert(passwordlessLogin)]: Can't insert passwordlessLogin. ")
	}
	store.loggerInfo.Logger.Infof("Passwordless_login_postgres_store: Insert - passwordles login is saved in database")

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
