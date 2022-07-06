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
	DATABASE   = "userData"
	COLLECTION = "userData"
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
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserDataMongoDBStore) GetByUserToken(token string) (*domain.UserData, error) {
	filter := bson.M{"token": token}
	return store.filterOne(filter)
}

func (store *UserDataMongoDBStore) Insert(userData *domain.UserData) error {
	_, err := store.userData.InsertOne(context.TODO(), userData)
	store.loggerInfo.Logger.Infof("user_data_job_mongodb_store: USCID | UI " + strconv.Itoa(userData.UserId))
	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: UFTSCIDD | UI " + strconv.Itoa(userData.UserId))
		return err
	}
	//userConnection.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserDataMongoDBStore) AddToken(userData *domain.UserData, newToken string) {
	//filter := bson.M{"user_id": id}
	//return store.filterOne(filter)
}

func (store *UserDataMongoDBStore) filter(filter interface{}) ([]*domain.UserData, error) {
	cursor, err := store.userData.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserDataMongoDBStore) filterOne(filter interface{}) (UserConnection *domain.UserData, err error) {
	result := store.userData.FindOne(context.TODO(), filter)
	err = result.Decode(&UserConnection)
	return
}

func decode(cursor *mongo.Cursor) (userData []*domain.UserData, err error) {
	for cursor.Next(context.TODO()) {
		var UserData domain.UserData
		err = cursor.Decode(&UserData)
		if err != nil {
			return
		}
		userData = append(userData, &UserData)
	}
	err = cursor.Err()
	return
}
