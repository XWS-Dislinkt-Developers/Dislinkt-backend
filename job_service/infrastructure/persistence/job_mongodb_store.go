package persistence

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

const (
	JOBDATABASE   = "jobData"
	JOBCOLLECTION = "jobData"
)

type JobMongoDBStore struct {
	jobData     *mongo.Collection
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewJobMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.JobOfferStore {
	jobData := client.Database(JOBDATABASE).Collection(JOBCOLLECTION)
	return &JobMongoDBStore{
		jobData:     jobData,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (store *JobMongoDBStore) GetByUserId(id int) (*domain.JobOffer, error) {
	filter := bson.M{"user_id": id}
	return store.filterOne(filter)
}

func (store *JobMongoDBStore) GetAll() ([]*domain.JobOffer, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *JobMongoDBStore) Insert(jobData *domain.JobOffer) error {
	_, err := store.jobData.InsertOne(context.TODO(), jobData)
	store.loggerInfo.Logger.Infof("Job_store_mongodb_store: USCID | UI " + strconv.Itoa(jobData.UserId))
	if err != nil {
		store.loggerError.Logger.Errorf("Job_store_mongodb_store: UFTSCIDD | UI " + strconv.Itoa(jobData.UserId))
		return err
	}
	//userConnection.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *JobMongoDBStore) filter(filter interface{}) ([]*domain.JobOffer, error) {
	cursor, err := store.jobData.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodej(cursor)
}

func (store *JobMongoDBStore) filterOne(filter interface{}) (JobOffer *domain.JobOffer, err error) {
	result := store.jobData.FindOne(context.TODO(), filter)
	err = result.Decode(&JobOffer)
	return
}

func decodej(cursor *mongo.Cursor) (jobData []*domain.JobOffer, err error) {
	for cursor.Next(context.TODO()) {
		var JobOffer domain.JobOffer
		err = cursor.Decode(&JobOffer)
		if err != nil {
			return
		}
		jobData = append(jobData, &JobOffer)
	}
	err = cursor.Err()
	return
}
