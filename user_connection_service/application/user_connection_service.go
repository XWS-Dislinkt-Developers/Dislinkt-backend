package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
	"time"
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

func (service *UserConnectionService) RegisterUserConnection(connection *domain.UserConnection) {
	err := service.store.Insert(connection)
	if err != nil {
		service.loggerError.Logger.Error("User_connection_service: CNSU ")
		println("Error in create method")
	}
}

func (service *UserConnectionService) Follow(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	if service.connectionDoesntExist(LoggedUserConnection, UserConnection) && service.requestDoesntExist(LoggedUserConnection, UserConnection) && service.BlockedDoesntExist(UserConnection, LoggedUserConnection) {

		if !service.requestDoesntExist(UserConnection, LoggedUserConnection) {
			UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
			LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
			service.store.UpdateConnections(UserConnection, LoggedUserConnection)
			var temp = domain.Notification{
				UserId:    idUser,
				Content:   "You have new connection",
				CreatedAt: time.Time{},
				Seen:      false,
			}
			var temp2 = domain.Notification{
				UserId:    idLoggedUser,
				Content:   "You have new connection",
				CreatedAt: time.Time{},
				Seen:      false,
			}
			service.store.InsertNotification(&temp)
			service.store.InsertNotification(&temp2)
		} else if UserConnection.Private {
			UserConnection.Requests = append(UserConnection.Requests, idLoggedUser)
			LoggedUserConnection.WaitingResponse = append(LoggedUserConnection.WaitingResponse, idUser)
			service.store.UpdateRequestConnection(UserConnection)
			service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
			var temp = domain.Notification{
				UserId:    idUser,
				Content:   "You have new request for connection",
				CreatedAt: time.Time{},
				Seen:      false,
			}

			service.store.InsertNotification(&temp)

		} else {
			if !service.waitingResponseDoesntExist(UserConnection, LoggedUserConnection) || !service.waitingResponseDoesntExist(LoggedUserConnection, UserConnection) {
				UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
				LoggedUserConnection.WaitingResponse = findAndDelete(LoggedUserConnection.WaitingResponse, idUser)
				service.store.UpdateWaitingResponseConnection(UserConnection)
				service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
			}
			if !service.requestDoesntExist(UserConnection, LoggedUserConnection) || !service.requestDoesntExist(LoggedUserConnection, UserConnection) {
				UserConnection.Requests = findAndDelete(UserConnection.Requests, idLoggedUser)
				LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
				service.store.UpdateRequestConnection(UserConnection)
				service.store.UpdateRequestConnection(LoggedUserConnection)
			}
			UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
			LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
			service.store.UpdateConnections(UserConnection, LoggedUserConnection)
		}
	}
}
func (service *UserConnectionService) Unfollow(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	UserConnection.Connections = findAndDelete(UserConnection.Connections, idLoggedUser)
	LoggedUserConnection.Connections = findAndDelete(LoggedUserConnection.Connections, idUser)
	service.store.UpdateConnections(UserConnection, LoggedUserConnection)
}
func (service *UserConnectionService) AcceptConnectionRequest(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
	service.store.UpdateRequestConnection(LoggedUserConnection)
	UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
	service.store.UpdateWaitingResponseConnection(UserConnection)

	UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
	LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
	service.store.UpdateConnections(UserConnection, LoggedUserConnection)

	var temp = domain.Notification{
		UserId:    idUser,
		Content:   "You request for connection is accepted",
		CreatedAt: time.Time{},
		Seen:      false,
	}

	service.store.InsertNotification(&temp)

}
func (service *UserConnectionService) DeclineConnectionRequest(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
	service.store.UpdateRequestConnection(LoggedUserConnection)
	UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
	service.store.UpdateWaitingResponseConnection(UserConnection)

}
func (service *UserConnectionService) CancelConnectionRequest(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	LoggedUserConnection.WaitingResponse = findAndDelete(LoggedUserConnection.WaitingResponse, idUser)
	service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
	UserConnection.Requests = findAndDelete(UserConnection.Requests, idLoggedUser)
	service.store.UpdateRequestConnection(UserConnection)

}
func (service *UserConnectionService) BlockUser(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	if service.BlockedDoesntExist(LoggedUserConnection, UserConnection) {
		if !service.connectionDoesntExist(LoggedUserConnection, UserConnection) {
			LoggedUserConnection.Connections = findAndDelete(LoggedUserConnection.Connections, idUser)
			UserConnection.Connections = findAndDelete(UserConnection.Connections, idLoggedUser)
			service.store.UpdateConnections(UserConnection, LoggedUserConnection)
		}
		if !service.requestDoesntExistForBlockingUsers(LoggedUserConnection, UserConnection) {
			LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
			service.store.UpdateRequestConnection(LoggedUserConnection)
			UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
			service.store.UpdateWaitingResponseConnection(UserConnection)
		}
		if !service.waitingResponseDoesntExistForBlockingUsers(LoggedUserConnection, UserConnection) {
			LoggedUserConnection.WaitingResponse = findAndDelete(LoggedUserConnection.WaitingResponse, idUser)
			service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
			UserConnection.Requests = findAndDelete(UserConnection.Requests, idLoggedUser)
			service.store.UpdateRequestConnection(UserConnection)
		}
		LoggedUserConnection.Blocked = append(LoggedUserConnection.Blocked, idUser)
		service.store.UpdateBlockedConnection(LoggedUserConnection)
	}
}
func (service *UserConnectionService) UnblockUser(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)

	LoggedUserConnection.Blocked = findAndDelete(LoggedUserConnection.Blocked, idUser)
	service.store.UpdateBlockedConnection(LoggedUserConnection)
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
func (service *UserConnectionService) requestDoesntExistForBlockingUsers(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range LoggedUserConnection.Requests {
		if c == UserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) waitingResponseDoesntExist(UserConnection *domain.UserConnection, LoggedUserConnection *domain.UserConnection) bool {
	for _, c := range UserConnection.Requests {
		if c == LoggedUserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) waitingResponseDoesntExistForBlockingUsers(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range LoggedUserConnection.WaitingResponse {
		if c == UserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) BlockedDoesntExist(loggedUser *domain.UserConnection, user *domain.UserConnection) bool {
	for _, c := range loggedUser.Blocked {
		if c == user.UserId {
			return false
		}
	}
	return true
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
