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
	DATABASE   = "messages"
	COLLECTION = "messages"
)

type MessagesMongoDBStore struct {
	messagesStore *mongo.Collection
	loggerInfo    *logg.Logger
	loggerError   *logg.Logger
}

func NewMessagesMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.MessageStore {
	messagesStore := client.Database(DATABASE).Collection(COLLECTION)
	return &MessagesMongoDBStore{
		messagesStore: messagesStore,
		loggerInfo:    loggerInfo,
		loggerError:   loggerError,
	}
}

func (store *MessagesMongoDBStore) GetBySenderId(id int) (*domain.Message, error) {
	filter := bson.M{"user_id": id}
	return store.filterOne(filter)
}

func (store *MessagesMongoDBStore) GetByReceiverId(id int) (*domain.Message, error) {
	filter := bson.M{"receiver_id": id}
	return store.filterOne(filter)
}

func (store *MessagesMongoDBStore) GetAll() ([]*domain.Message, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *MessagesMongoDBStore) Insert(mess *domain.Message) error {
	_, err := store.messagesStore.InsertOne(context.TODO(), mess)
	store.loggerInfo.Logger.Infof("Message_mongodb_store: USCID | UI " + strconv.Itoa(mess.SenderId))
	if err != nil {
		store.loggerError.Logger.Errorf("User_connection_mongodb_store: UFTSCIDD | UI " + strconv.Itoa(mess.SenderId))
		return err
	}
	//userConnection.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *MessagesMongoDBStore) filter(filter interface{}) ([]*domain.Message, error) {
	cursor, err := store.messagesStore.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *MessagesMongoDBStore) filterOne(filter interface{}) (UserConnection *domain.Message, err error) {
	result := store.messagesStore.FindOne(context.TODO(), filter)
	err = result.Decode(&UserConnection)
	return
}

func decode(cursor *mongo.Cursor) (userConnections []*domain.Message, err error) {
	for cursor.Next(context.TODO()) {
		var UserConnection domain.Message
		err = cursor.Decode(&UserConnection)
		if err != nil {
			return
		}
		userConnections = append(userConnections, &UserConnection)
	}
	err = cursor.Err()
	return
}
