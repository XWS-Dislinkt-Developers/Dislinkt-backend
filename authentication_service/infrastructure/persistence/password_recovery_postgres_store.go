package persistence

import (
	"errors"
	ftm "fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	"gorm.io/gorm"
	"strconv"
)

type PasswordRecoveryPostgresStore struct {
	db          *gorm.DB
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewPasswordRecoveryPostgresStore(db *gorm.DB, loggerInfo *logg.Logger, loggerError *logg.Logger) (domain.PasswordRecoveryStore, error) {
	err := db.AutoMigrate(&domain.PasswordRecovery{})
	if err != nil {
		loggerError.Logger.Errorf("Password_recovery_postgres_store: - failed method   ")
		return nil, err
	}
	return &PasswordRecoveryPostgresStore{db: db,
		loggerInfo:  loggerInfo,
		loggerError: loggerError}, nil
}

func (store *PasswordRecoveryPostgresStore) GetByRecoveryCode(recoveryCode string) (passwordRecovery *domain.PasswordRecovery, err error) {
	var password *domain.PasswordRecovery
	passwords, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Password_recovery_postgres_store: NUWC ")
		return nil, errors.New("[PasswordRecoveryPostgresStore-GetByRecoveryCode(confToken)]: There's no user.")
	}
	for _, p := range *passwords {
		if p.RecoveryCode == recoveryCode {
			return &p, nil
		}
	}
	if password == nil {
		store.loggerError.Logger.Errorf("Password_recovery_postgres_store: NUWC ")

		return nil, errors.New("ERR - [ConfirmationTokenPostgresStore-GetByConfirmationToken(confToken)]: Can't find user with this token: " + recoveryCode)
	}
	store.loggerInfo.Logger.Infof("Password_recovery_postgres_store: FUWRC ")
	return password, nil
}

func (store *PasswordRecoveryPostgresStore) GetByUserId(id int) (*domain.PasswordRecovery, error) {
	var password *domain.PasswordRecovery
	passwords, err := store.GetAll()
	if err != nil {
		store.loggerError.Logger.Errorf("Password_recovery_postgres_store: NUUI  " + strconv.Itoa(id))
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
		store.loggerError.Logger.Errorf("Password_recovery_postgres_store: PSWRCNSD!")
		ftm.Println("[PasswordRecoveryPostgresStore-Insert(passwordRecovery)]: Can't insert passwordRecovery.")
		return errors.New("ERR - [PasswordRecoveryPostgresStore-Insert(passwordRecovery)]: Can't insert passwordRecovery. ")
	}
	store.loggerInfo.Logger.Infof("Password_recovery_postgres_store: NPSWRS ")

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
	passwordRecovery, _ := store.GetByUserId(idUser)
	store.db.Where("user_id = ?", idUser).Delete(&passwordRecovery)
}
