package application

import "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"

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

func (service *UserService) Get(id string) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	return service.store.GetByUsername(username)
}

func (service *UserService) GetAll() (*[]domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) DeleteAll() {
	service.store.DeleteAll()
}
