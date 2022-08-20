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

func NewJobService(userstore domain.UserDataStore, jobstore domain.JobOfferStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *JobService {
	return &JobService{
		jobstore:    jobstore,
		userstore:   userstore,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *JobService) GetAllUserData() ([]*domain.UserData, error) {
	return service.userstore.GetAll()
}

func (service *JobService) GetUserDataById(idUser int) (*domain.UserData, error) {
	return service.userstore.GetByUserId(idUser)
}

func (service *JobService) GetByUserToken(token string) (*domain.UserData, error) {
	return service.userstore.GetByUserToken(token)
}

func (service *JobService) InsertUserData(userData *domain.UserData) error {
	return service.userstore.Insert(userData)
}

func (service *JobService) AddToken(id int64) (string, error) {

	//domain.UserData{UserId: id,Id: 1,Token: "newToken"} data
	//err := service.InsertUserData(data)
	return "", nil
}

func (service *JobService) GetAllJobData() ([]*domain.JobOffer, error) {
	return service.jobstore.GetAll()
}

func (service *JobService) GetJobDataByCompany(company string) ([]*domain.JobOffer, error) {
	return service.jobstore.GetByCompany(company)
}

func (service *JobService) GetJobDataById(idUser int) (*domain.JobOffer, error) {
	return service.jobstore.GetByUserId(idUser)
}

func (service *JobService) InsertJobData(jobData *domain.JobOffer) error {
	return service.jobstore.Insert(jobData)
}
