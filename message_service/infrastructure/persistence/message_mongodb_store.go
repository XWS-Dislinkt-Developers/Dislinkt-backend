package persistence

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

const (
	DATABASE     = "userMessages"
	COLLECTION   = "userMessages"
	MNDATABASE   = "userMessageNotifications"
	MNCOLLECTION = "userMessageNotifications"
)

type MessagesMongoDBStore struct {
	messagesStore             *mongo.Collection
	notificationMessagesStore *mongo.Collection
	loggerInfo                *logg.Logger
	loggerError               *logg.Logger
}

func NewMessagesMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.MessageStore {
	messagesStore := client.Database(DATABASE).Collection(COLLECTION)
	notificationMessagesStore := client.Database(MNDATABASE).Collection(MNCOLLECTION)
	return &MessagesMongoDBStore{
		messagesStore:             messagesStore,
		notificationMessagesStore: notificationMessagesStore,
		loggerInfo:                loggerInfo,
		loggerError:               loggerError,
	}
}

// CRUD - READ method(s) - GetAll, GetAllUsersMessagesByUserId
func (store *MessagesMongoDBStore) GetAll() ([]*domain.Message, error) {
	noFiltering := bson.D{{}}
	return store.filter(noFiltering)
}
func (store *MessagesMongoDBStore) GetAllUsersMessagesByUserId(id int) ([]*domain.Message, error) {
	filteringUserMessages := bson.M{"$or": []bson.M{{"sender_id": id}, {"receiver_id": id}}}
	return store.filter(filteringUserMessages)
}

/*
func (store *MessagesMongoDBStore) GetAllReceiversMessagesByUserId(id int) ([]*domain.Message, error) {
	filteringReceiversMessages := bson.M{"receiver_id": id}
	return store.filter(filteringReceiversMessages)
}
func (store *MessagesMongoDBStore) GetAllMessagesBetweenUsers(userId1, userId2 int) ([]*domain.Message, error) {
	filteringMessagesBetweenUsers := bson.M{"$or": []bson.M{{"sender_id": userId1, "receiver_id": userId2}, {"sender_id": userId2, "receiver_id": userId1}}}
	return store.filter(filteringMessagesBetweenUsers)
}
*/

// CRUD - CREATE method(s) - Insert
func (store *MessagesMongoDBStore) Insert(message *domain.Message) error {
	_, err := store.messagesStore.InsertOne(context.TODO(), message)
	store.loggerInfo.Logger.Infof("Message_mongodb_store: USCID | UI " + strconv.Itoa(message.SenderId))
	if err != nil {
		store.loggerError.Logger.Errorf("Message_mongodb_store: UFTSCIDD | UI " + strconv.Itoa(message.SenderId))
		return err
	}
	//message.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

// CRUD - DELETE method(s) - DeleteAll
func (store *MessagesMongoDBStore) DeleteAll() {
	store.messagesStore.DeleteMany(context.TODO(), bson.D{{}})
}

// Helper method(s) - filter, filterOne, decode
func (store *MessagesMongoDBStore) filter(filter interface{}) ([]*domain.Message, error) {
	cursor, err := store.messagesStore.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *MessagesMongoDBStore) filterNotification(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notificationMessagesStore.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	return decodeNotification(cursor)
}

func (store *MessagesMongoDBStore) filterOne(filter interface{}) (Message *domain.Message, err error) {
	result := store.messagesStore.FindOne(context.TODO(), filter)
	err = result.Decode(&Message)
	return
}
func decode(cursor *mongo.Cursor) (Messages []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var Message domain.Message
		err = cursor.Decode(&Message)
		if err != nil {
			return
		}
		Messages = append(Messages, &Message)
	}
	err = cursor.Err()
	return
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

func (store *MessagesMongoDBStore) InsertNotification(notification *domain.Notification) error {
	_, err := store.notificationMessagesStore.InsertOne(context.TODO(), notification)
	store.loggerInfo.Logger.Infof("Message_notification_mongodb_store: USCID | UI " + strconv.Itoa(notification.UserId))
	if err != nil {
		println("erorr while inserting notification")
		return err
	}
	//message.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *MessagesMongoDBStore) GetAllUserNotificationsByUserId(id int) ([]*domain.Notification, error) {
	filteringUserNotifications := bson.M{"$or": []bson.M{{"user_id": id}}}
	return store.filterNotification(filteringUserNotifications)
}
