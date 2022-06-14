package persistence

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

const (
	DATABASE   = "userConnection"
	COLLECTION = "userConnection"
)

type UserConnectionMongoDBStore struct {
	userConnections *mongo.Collection
	loggerInfo      *logg.Logger
	loggerError     *logg.Logger
}

func NewUserConnectionMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.UserConnectionStore {
	userConnections := client.Database(DATABASE).Collection(COLLECTION)
	return &UserConnectionMongoDBStore{
		userConnections: userConnections,
		loggerInfo:      loggerInfo,
		loggerError:     loggerError,
	}
}

func (store *UserConnectionMongoDBStore) GetByUserId(id int) (*domain.UserConnection, error) {
	filter := bson.M{"user_id": id}
	return store.filterOne(filter)
}

func (store *UserConnectionMongoDBStore) GetAll() ([]*domain.UserConnection, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserConnectionMongoDBStore) Insert(userConnection *domain.UserConnection) error {
	_, err := store.userConnections.InsertOne(context.TODO(), userConnection)
	store.loggerInfo.Logger.Infof("User_connection_mongodb_store: Insert - User with id " + strconv.Itoa(userConnection.UserId) + " save his connection in database ")
	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: Insert - failed method - couldn't save user connection in database")
		return err
	}
	//userConnection.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserConnectionMongoDBStore) DeleteAll() {
	store.userConnections.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserConnectionMongoDBStore) filter(filter interface{}) ([]*domain.UserConnection, error) {
	cursor, err := store.userConnections.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *UserConnectionMongoDBStore) filterOne(filter interface{}) (UserConnection *domain.UserConnection, err error) {
	result := store.userConnections.FindOne(context.TODO(), filter)
	err = result.Decode(&UserConnection)
	return
}

func (store *UserConnectionMongoDBStore) UpdateRequestConnection(userConnection *domain.UserConnection) {
	_, err := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": userConnection.UserId}, bson.D{{"$set", bson.D{{"requests", userConnection.Requests}}}})

	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: UpdateRequestConnection - failed method - User with id " + strconv.Itoa(userConnection.UserId) + " couldn't update his connection in database ")

		println("Failed update request connection.")
	} else {
		store.loggerInfo.Logger.Infof("User_connection_mongodb_store: UpdaterequestConnection - User with id " + strconv.Itoa(userConnection.UserId) + " updated his connections in database")

	}
}

func (store *UserConnectionMongoDBStore) AddConnections(userConnection *domain.UserConnection, loggedUserConnection *domain.UserConnection) {
	_, err1 := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": userConnection.UserId}, bson.D{{"$set", bson.D{{"connections", userConnection.Connections}}}})
	_, err2 := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": loggedUserConnection.UserId}, bson.D{{"$set", bson.D{{"connections", loggedUserConnection.Connections}}}})
	if err1 != nil || err2 != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: AddConnections - failed method - User with id " + strconv.Itoa(loggedUserConnection.UserId) + " couldn't add connection in database ")

		println("Failed update connection.")
	} else {
		store.loggerInfo.Logger.Infof("User_connection_mongodb_store: AddConnections - User with id " + strconv.Itoa(loggedUserConnection.UserId) + " save connection in database")

	}
}

func decode(cursor *mongo.Cursor) (userConnections []*domain.UserConnection, err error) {
	for cursor.Next(context.TODO()) {
		var UserConnection domain.UserConnection
		err = cursor.Decode(&UserConnection)
		if err != nil {
			return
		}
		userConnections = append(userConnections, &UserConnection)
	}
	err = cursor.Err()
	return
}
