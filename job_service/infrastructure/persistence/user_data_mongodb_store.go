package persistence

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE      = "userData"
	COLLECTION    = "userData"
	JOBDATABASE   = "jobData"
	JOBCOLLECTION = "jobData"
)

type UserDataMongoDBStore struct {
	userData    *mongo.Collection
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserDataMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.UserDataStore {
	userData := client.Database(DATABASE).Collection(COLLECTION)
	return &UserDataMongoDBStore{
		userData:    userData,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (store *UserDataMongoDBStore) GetByUserId(id int) (*domain.UserData, error) {
	filter := bson.M{"user_id": id}
	return store.filterOne(filter)
}

func (store *UserDataMongoDBStore) GetAll() ([]*domain.UserData, error) {
	//filter := bson.D{{}}
	//return store.filter(filter)
	return nil, nil
}

func (store *UserDataMongoDBStore) GetByUserToken(token string) (*domain.UserData, error) {
	//filter := bson.M{"user_id": id}
	//return store.filterOne(filter)
	return nil, nil
}

func (store *UserDataMongoDBStore) Insert(userData *UserData) (*domain.UserData, error) {
	//filter := bson.M{"user_id": id}
	//return store.filterOne(filter)
	return nil, nil
}

func (store *UserDataMongoDBStore) AddToken(userData *UserData, newToken string) (*domain.UserData, error) {
	//filter := bson.M{"user_id": id}
	//return store.filterOne(filter)
	return nil, nil
}
