package persistence

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "userConnection"
	COLLECTION = "userConnection"
)

type UserConnectionMongoDBStore struct {
	userConnections *mongo.Collection
}

func NewUserConnectionMongoDBStore(client *mongo.Client) domain.UserConnectionStore {
	userConnections := client.Database(DATABASE).Collection(COLLECTION)
	return &UserConnectionMongoDBStore{
		userConnections: userConnections,
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
	if err != nil {
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
