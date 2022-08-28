package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"
)

type NotificationService struct {
	store       domain.MessageStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewNotificationService(store domain.MessageStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *NotificationService {
	return &NotificationService{
		store:       store,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *NotificationService) GetAllUserNotificationsByUserId(idUser int) ([]*domain.Notification, error) {
	return service.store.GetAllUserNotificationsByUserId(idUser)
}

func (service *NotificationService) InsertNotification(notification *domain.Notification) {
	err := service.store.InsertNotification(notification)
	if err != nil {
		service.loggerError.Logger.Error("Message_notification_service: CNSU ")
		println("Error in insert notification method")
	}
}
