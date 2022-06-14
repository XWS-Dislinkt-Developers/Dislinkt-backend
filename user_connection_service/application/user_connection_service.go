package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
)

type UserConnectionService struct {
	store       domain.UserConnectionStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserConnectionService(store domain.UserConnectionStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserConnectionService {
	return &UserConnectionService{
		store:       store,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *UserConnectionService) GetAll() ([]*domain.UserConnection, error) {
	return service.store.GetAll()
}

func (service *UserConnectionService) GetConnectionsById(idUser int) (*domain.UserConnection, error) {
	return service.store.GetByUserId(idUser)
}

//user je korisnik sa kojim ulogovani korisnik zeli da ostvari konekciju
func (service *UserConnectionService) Follow(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	if service.connectionDoesntExist(LoggedUserConnection, UserConnection) && service.requestDoesntExist(LoggedUserConnection, UserConnection) {
		if UserConnection.Private {
			UserConnection.Requests = append(UserConnection.Requests, idLoggedUser)
			service.store.UpdateRequestConnection(UserConnection)
		} else {
			UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
			LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
			service.store.AddConnections(UserConnection, LoggedUserConnection)
		}
	}
}

func (service *UserConnectionService) connectionDoesntExist(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range LoggedUserConnection.Connections {
		if c == UserConnection.UserId {
			return false
		}
	}
	return true
}

func (service *UserConnectionService) requestDoesntExist(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range UserConnection.Requests {
		if c == LoggedUserConnection.UserId {
			return false
		}
	}
	return true
}

func (service *UserConnectionService) Unfollow(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	UserConnection.Connections = findAndDelete(UserConnection.Connections, idLoggedUser)
	LoggedUserConnection.Connections = findAndDelete(LoggedUserConnection.Connections, idUser)
	service.store.AddConnections(UserConnection, LoggedUserConnection)
}

func (service *UserConnectionService) AcceptConnectionRequest(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
	service.store.UpdateRequestConnection(LoggedUserConnection)

	UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
	LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
	service.store.AddConnections(UserConnection, LoggedUserConnection)
}

func (service *UserConnectionService) DeclineConnectionRequest(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)

	LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
	service.store.UpdateRequestConnection(LoggedUserConnection)
}

func findAndDelete(s []int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}
