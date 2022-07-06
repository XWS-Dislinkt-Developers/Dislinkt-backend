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

func (service *MessageService) GetAll() ([]*domain.Message, error) {
	return service.store.GetAll()
}

func (service *MessageService) GetConnectionsById(idUser int) (*domain.Message, error) {
	return nil, nil
	//return service.store.GetByUserId(idUser), nil
}

func (service *MessageService) Insert(mess *domain.Message) error {
	return nil
}
