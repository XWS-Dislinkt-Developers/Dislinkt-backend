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

//user je korisnik sa kojim ulogovani korisnik zeli da ostvari konekciju
func (service *UserConnectionService) Follow(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	if UserConnection.Private {
		UserConnection.Requests = append(UserConnection.Requests, idLoggedUser)
		service.store.AddRequestConnection(UserConnection)
	} else {
		UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
		LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
		service.store.AddConnections(UserConnection, LoggedUserConnection)
	}
}
