package application

import (
	"errors"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
)

type UserService struct {
	store                  domain.UserStore
	conformationTokenStore domain.ConfirmationTokenStore
	loggerInfo             *logg.Logger
	loggerError            *logg.Logger
}

func NewUserService(store domain.UserStore, conformationTokenStore domain.ConfirmationTokenStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserService {
	return &UserService{
		store:                  store,
		conformationTokenStore: conformationTokenStore,
		loggerInfo:             loggerInfo,
		loggerError:            loggerError,
	}
}

func (service *UserService) Create(user *domain.User) {
	err := service.store.Insert(user)
	if err != nil {
		println("Error in create method")
	}

}

func (service *UserService) Get(id int) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	return service.store.GetByUsername(username)
}

func (service *UserService) UpdateUser(dto domain.UpdateUserDto, userID int) (*domain.User, error) {
	foundUser, _ := service.GetByUsername(dto.Username)
	if foundUser != nil && foundUser.ID != userID {
		return nil, errors.New("Username is already taken")
	}

	return service.store.UpdateUser(dto, userID)
}

func (service *UserService) UpdateUserWAE(dto domain.UpdateUserWAEDto, userID int) (*domain.User, error) {
	return service.store.UpdateUserWorkAndEducation(dto, userID)
}
func (service *UserService) UpdateUserSAI(dto domain.UpdateUserSAIDto, userID int) (*domain.User, error) {
	return service.store.UpdateUserSkillsAndInterests(dto, userID)
}

func (service *UserService) GetAll() (*[]domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) DeleteAll() {
	service.store.DeleteAll()
}
