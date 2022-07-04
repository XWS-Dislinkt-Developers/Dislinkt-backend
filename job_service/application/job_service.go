package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/logger"
)

type JobService struct {
	userstore   domain.UserDataStore
	jobstore    domain.JobOfferStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewJobService(userstore domain.UserDataStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *JobService {
	return &JobService{
		userstore:   userstore,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *JobService) GetAll() ([]*domain.UserData, error) {
	return service.userstore.GetAll()
}

func (service *JobService) GetUserDataById(idUser int) (*domain.UserData, error) {
	return service.userstore.GetByUserId(idUser)
}
