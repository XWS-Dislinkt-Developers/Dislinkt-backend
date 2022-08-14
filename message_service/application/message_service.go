package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"
)

type MessageService struct {
	store       domain.MessageStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewMessageService(store domain.MessageStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *MessageService {
	return &MessageService{
		store:       store,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

// CRUD - READ method(s)
func (service *MessageService) GetAll() ([]*domain.Message, error) {
	return service.store.GetAll()
}
func (service *MessageService) GetAllUsersMessagesByUserId(idUser int) ([]*domain.Message, error) {
	return service.store.GetAllUsersMessagesByUserId(idUser)
}

/*
func (service *MessageService) GetAllReceiversMessagesByUserId(idUser int) ([]*domain.Message, error) {
	return service.store.GetAllReceiversMessagesByUserId(idUser)
}
func (service *MessageService) GetAllMessagesBetweenUsers(userId1, userId2 int) ([]*domain.Message, error) {
	return service.store.GetAllMessagesBetweenUsers(userId1, userId2)
}
*/

// CRUD - CREATE method(s)
func (service *MessageService) Insert(message *domain.Message) {
	err := service.store.Insert(message)
	if err != nil {
		service.loggerError.Logger.Error("Message_service: CNSU ")
		println("Error in create method")
	}
}
