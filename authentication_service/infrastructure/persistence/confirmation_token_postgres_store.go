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
		loggerError.Logger.Errorf("Confirmation_token_postgres_store: - failed method   ")

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
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: NUUI  " + strconv.Itoa(id))

		return nil, errors.New("[ConfirmationTokenPostgresStore-GetByUserId(id)]: There's no user.")
	}
	for _, t := range *tokens {
		if t.UserId == id {
			return &t, nil
		}
	}
	if token == nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: TIN ")
		return nil, errors.New("ERR - [ConfirmationTokenPostgresStore-GetByUserId(id)]: Can't find user with this id: " + string(id))
	}
	return token, nil
}

func (store *ConfirmationTokenPostgresStore) GetByConfirmationToken(confToken string) (*domain.ConfirmationToken, error) {
	var token *domain.ConfirmationToken
	tokens, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: NT")
		return nil, errors.New("[ConfirmationTokenPostgresStore-GetByConfirmationToken(confToken)]: There's no token.")
	}
	for _, t := range *tokens {
		if t.ConfirmationToken == confToken {
			return &t, nil
		}
	}
	if token == nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: CNFUT ")
		return nil, errors.New("ERR - [ConfirmationTokenPostgresStore-GetByConfirmationToken(confToken)]: Can't find user with this token: " + confToken)
	}
	return token, nil
}

func (store *ConfirmationTokenPostgresStore) Insert(confirmationToken *domain.ConfirmationToken) error {
	result := store.db.Create(confirmationToken)
	if result.Error != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: CNSCTD ")

		ftm.Println("[ConfirmationTokenPostgresStore-Insert(confirmationToken)]: Can't insert confirmationToken.")
		return errors.New("ERR - [ConfirmationTokenPostgresStore-Insert(confirmationToken)]: Can't insert confirmationToken. ")
	}
	store.loggerInfo.Logger.Infof("Confirmation_token_postgres_store: CTSD ")
	return nil
}

func (store *ConfirmationTokenPostgresStore) GetAll() (*[]domain.ConfirmationToken, error) {
	var tokens []domain.ConfirmationToken
	result := store.db.Find(&tokens)
	if result.Error != nil {
		store.loggerError.Logger.Errorf("Confirmation_token_postgres_store: CNRDD ! ")

		return nil, result.Error
	}
	return &tokens, nil
}

func (store *ConfirmationTokenPostgresStore) Delete(idUser int) {
	token, _ := store.GetByUserId(idUser)
	store.db.Where("user_id = ?", idUser).Delete(&token)
	store.loggerInfo.Logger.Infof("Confirmation_token_postgres_store: Delete - CTDFD | UI " + strconv.Itoa(idUser))

}
