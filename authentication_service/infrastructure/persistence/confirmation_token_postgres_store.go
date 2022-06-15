package persistence

import (
	"errors"
	ftm "fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	"gorm.io/gorm"
	"strconv"
)

type ConfirmationTokenPostgresStore struct {
	db          *gorm.DB
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewConfirmationTokenPostgresStore(db *gorm.DB, loggerInfo *logg.Logger, loggerError *logg.Logger) (domain.ConfirmationTokenStore, error) {
	err := db.AutoMigrate(&domain.ConfirmationToken{})
	if err != nil {
		loggerError.Logger.Errorf("Confirmation_token_postgres_store: NewConfirmationTokenPostgresStore - failed method   ")

		return nil, err
	}
	return &ConfirmationTokenPostgresStore{db: db,
		loggerInfo:  loggerInfo,
		loggerError: loggerError}, nil
}

func (store *ConfirmationTokenPostgresStore) GetByUserId(id int) (*domain.ConfirmationToken, error) {
	var token *domain.ConfirmationToken
	tokens, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: GetByUserId - failed method - there is no user with user id " + strconv.Itoa(id))

		return nil, errors.New("[ConfirmationTokenPostgresStore-GetByUserId(id)]: There's no user.")
	}
	for _, t := range *tokens {
		if t.UserId == id {
			return &t, nil
		}
	}
	if token == nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: GetByUserId - failed method - token is null ")
		return nil, errors.New("ERR - [ConfirmationTokenPostgresStore-GetByUserId(id)]: Can't find user with this id: " + string(id))
	}
	return token, nil
}

func (store *ConfirmationTokenPostgresStore) GetByConfirmationToken(confToken string) (*domain.ConfirmationToken, error) {
	var token *domain.ConfirmationToken
	tokens, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: GetByConfirmationToken - failed method - there is no token ")
		return nil, errors.New("[ConfirmationTokenPostgresStore-GetByConfirmationToken(confToken)]: There's no token.")
	}
	for _, t := range *tokens {
		if t.ConfirmationToken == confToken {
			return &t, nil
		}
	}
	if token == nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: GetByConfirmationToken - failed method - can't find user with this token ")
		return nil, errors.New("ERR - [ConfirmationTokenPostgresStore-GetByConfirmationToken(confToken)]: Can't find user with this token: " + confToken)
	}
	return token, nil
}

func (store *ConfirmationTokenPostgresStore) Insert(confirmationToken *domain.ConfirmationToken) error {
	result := store.db.Create(confirmationToken)
	if result.Error != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: Insert - failed method - can't save confirmation token in database ")

		ftm.Println("[ConfirmationTokenPostgresStore-Insert(confirmationToken)]: Can't insert confirmationToken.")
		return errors.New("ERR - [ConfirmationTokenPostgresStore-Insert(confirmationToken)]: Can't insert confirmationToken. ")
	}
	store.loggerInfo.Logger.Infof("Confirmation_token_postgres_store: Insert - new confirmation token is saved in database")
	return nil
}

func (store *ConfirmationTokenPostgresStore) GetAll() (*[]domain.ConfirmationToken, error) {
	var tokens []domain.ConfirmationToken
	result := store.db.Find(&tokens)
	if result.Error != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: GetAll - failed method - can't read data from database! ")

		return nil, result.Error
	}
	return &tokens, nil
}

func (store *ConfirmationTokenPostgresStore) Delete(idUser int) {
	token, _ := store.GetByUserId(idUser)
	store.db.Where("user_id = ?", idUser).Delete(&token)
	store.loggerInfo.Logger.Infof("Confirmation_token_postgres_store: Delete - confirmation token is deleted from database for user " + strconv.Itoa(idUser))

}
