package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
)

type UserConnectionService struct {
	store domain.UserConnectionStore
}

func NewUserConnectionService(store domain.UserConnectionStore) *UserConnectionService {
	return &UserConnectionService{
		store: store,
	}
}

func (service *UserConnectionService) GetAll() ([]*domain.UserConnection, error) {
	return service.store.GetAll()
}
