package application

import (
	"errors"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
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

	return service.store.UpdateUser(dto)
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
