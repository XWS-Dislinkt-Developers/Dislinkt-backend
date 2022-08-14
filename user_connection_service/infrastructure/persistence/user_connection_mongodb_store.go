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
	DATABASE      = "userConnection"
	COLLECTION    = "userConnection"
	UCNDATABASE   = "userConnectionNotifications"
	UCNCOLLECTION = "userConnectionNotifications"
)

type UserConnectionMongoDBStore struct {
	userConnections           *mongo.Collection
	notificationMessagesStore *mongo.Collection
	loggerInfo                *logg.Logger
	loggerError               *logg.Logger
}

func NewUserConnectionMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.UserConnectionStore {
	userConnections := client.Database(DATABASE).Collection(COLLECTION)
	notificationMessagesStore := client.Database(UCNDATABASE).Collection(UCNCOLLECTION)
	return &UserConnectionMongoDBStore{
		userConnections:           userConnections,
		notificationMessagesStore: notificationMessagesStore,
		loggerInfo:                loggerInfo,
		loggerError:               loggerError,
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
	store.loggerInfo.Logger.Infof("User_connection_mongodb_store: USCID | UI " + strconv.Itoa(userConnection.UserId))
	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: UFTSCIDD | UI " + strconv.Itoa(userConnection.UserId))
		return err
	}
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
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: FTUCID | UI  " + strconv.Itoa(userConnection.UserId))

		println("Failed update request connection.")
	} else {
		store.loggerInfo.Logger.Infof("User_connection_mongodb_store: USUCID | UI " + strconv.Itoa(userConnection.UserId))

	}
}

func (store *UserConnectionMongoDBStore) UpdateWaitingResponseConnection(userConnection *domain.UserConnection) {
	_, err := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": userConnection.UserId}, bson.D{{"$set", bson.D{{"waiting_response", userConnection.WaitingResponse}}}})
	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: FTUCID | UI  " + strconv.Itoa(userConnection.UserId))

		println("Failed update request connection.")
	} else {
		store.loggerInfo.Logger.Infof("User_connection_mongodb_store: USUCID | UI " + strconv.Itoa(userConnection.UserId))

	}
}

func (store *UserConnectionMongoDBStore) UpdateBlockedConnection(userConnection *domain.UserConnection) {
	_, err := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": userConnection.UserId}, bson.D{{"$set", bson.D{{"blocked", userConnection.Blocked}}}})
	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: FTBCID | UI  " + strconv.Itoa(userConnection.UserId))
		println("Failed update blocked connection.")
	}
}

func (store *UserConnectionMongoDBStore) UpdateConnections(userConnection *domain.UserConnection, loggedUserConnection *domain.UserConnection) {
	_, err1 := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": userConnection.UserId}, bson.D{{"$set", bson.D{{"connections", userConnection.Connections}}}})
	_, err2 := store.userConnections.UpdateOne(context.TODO(), bson.M{"user_id": loggedUserConnection.UserId}, bson.D{{"$set", bson.D{{"connections", loggedUserConnection.Connections}}}})
	if err1 != nil || err2 != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: UFACID | UI " + strconv.Itoa(loggedUserConnection.UserId))
		println("Failed update connection.")
	} else {
		store.loggerInfo.Logger.Infof("User_connection_mongodb_store: USSACID | UI " + strconv.Itoa(loggedUserConnection.UserId))
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

func (store *UserConnectionMongoDBStore) filterNotification(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notificationMessagesStore.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return decodeNotification(cursor)
}

func decodeNotification(cursor *mongo.Cursor) (Notifications []*domain.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var Notification domain.Notification
		err = cursor.Decode(&Notification)
		if err != nil {
			return
		}
		Notifications = append(Notifications, &Notification)
	}
	err = cursor.Err()
	return
}

func (store *UserConnectionMongoDBStore) InsertNotification(notification *domain.Notification) error {
	_, err := store.notificationMessagesStore.InsertOne(context.TODO(), notification)
	store.loggerInfo.Logger.Infof("User_connection_notification_mongodb_store: USCID | UI " + strconv.Itoa(notification.UserId))
	if err != nil {
		println("erorr while inserting notification")
		return err
	}
	//message.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserConnectionMongoDBStore) GetAllUserNotificationsByUserId(id int) ([]*domain.Notification, error) {
	filteringUserNotifications := bson.M{"$or": []bson.M{{"user_id": id}}}
	return store.filterNotification(filteringUserNotifications)
}
