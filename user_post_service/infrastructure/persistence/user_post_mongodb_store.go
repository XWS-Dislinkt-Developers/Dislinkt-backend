package persistence

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

const (
	DATABASE     = "userPost"
	COLLECTION   = "userPost"
	PNDATABASE   = "userPostNotifications"
	PNCOLLECTION = "userPostNotifications"
)

type UserPostMongoDBStore struct {
	userPosts             *mongo.Collection
	notificationPostStore *mongo.Collection
	loggerInfo            *logg.Logger
	loggerError           *logg.Logger
}

func NewUserPostMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.UserPostStore {
	userPosts := client.Database(DATABASE).Collection(COLLECTION)
	notificationPostStore := client.Database(PNDATABASE).Collection(PNCOLLECTION)
	return &UserPostMongoDBStore{
		userPosts:             userPosts,
		notificationPostStore: notificationPostStore,
		loggerInfo:            loggerInfo,
		loggerError:           loggerError,
	}
}

func (store *UserPostMongoDBStore) Get(id primitive.ObjectID) (*domain.UserPost, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *UserPostMongoDBStore) GetAll() ([]*domain.UserPost, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *UserPostMongoDBStore) Insert(userPost *domain.UserPost) error {
	//userPost.Id = primitive.NewObjectID()
	result, err := store.userPosts.InsertOne(context.TODO(), userPost)
	if err != nil {
		store.loggerError.Logger.Errorf("User_post_mongodb_store: UFSNPD | UI  " + strconv.Itoa(userPost.UserId))
		return err
	}
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: USSNPD | UI  " + strconv.Itoa(userPost.UserId))
	userPost.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserPostMongoDBStore) DeleteAll() {
	store.userPosts.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *UserPostMongoDBStore) filter(filter interface{}) ([]*domain.UserPost, error) {
	cursor, err := store.userPosts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {

		return nil, err
	}
	return decode(cursor)
}

func (store *UserPostMongoDBStore) filterOne(filter interface{}) (UserPost *domain.UserPost, err error) {
	result := store.userPosts.FindOne(context.TODO(), filter)
	err = result.Decode(&UserPost)
	return
}

func (store *UserPostMongoDBStore) UpdateComments(userPost *domain.UserPost) {
	_, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"comments", userPost.Comments}}}})
	if err != nil {
		println("failed")
		store.loggerError.Logger.Errorf("User_post_mongodb_store: UFSNCOPID | UI  " + strconv.Itoa(userPost.UserId))
	}
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: USSNCOPID  | UI  " + strconv.Itoa(userPost.UserId))
	println("success update ")
}

func (store *UserPostMongoDBStore) AddReaction(userPost *domain.UserPost) {
	_, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"reactions", userPost.Reactions}}}})
	if err != nil {
		println("failed reaction update")
		store.loggerError.Logger.Errorf("User_post_mongodb_store: UFSNROPID | UI " + strconv.Itoa(userPost.UserId))
	}
	println("success reaction update ")
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: USSNROPID | UI " + strconv.Itoa(userPost.UserId))
}

func (store *UserPostMongoDBStore) UpdateReactions(userReaction *domain.Reaction, userPost *domain.UserPost) {
	_, err := store.userPosts.UpdateOne(
		context.TODO(),
		bson.M{"_id": userPost.Id, "reactions.user_id": userReaction.UserId},
		bson.D{
			{"$set", bson.M{"reactions.0.liked": userReaction.Liked, "reactions.0.disliked": userReaction.Disliked}},
		})
	//0 because that's first structure in array
	if err != nil {
		store.loggerError.Logger.Errorf("User_post_mongodb_store: UFSNROPID | UI  " + strconv.Itoa(userPost.UserId))

	}
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: USSNROPID | UI  " + strconv.Itoa(userPost.UserId))
}

func (store *UserPostMongoDBStore) GetPostsByUserId(userId int) ([]*domain.UserPost, error) {
	filter := bson.M{"user_id": userId}
	return store.filter(filter)
}

func decode(cursor *mongo.Cursor) (userPosts []*domain.UserPost, err error) {
	for cursor.Next(context.TODO()) {
		var UserPost domain.UserPost
		err = cursor.Decode(&UserPost)
		if err != nil {
			return
		}
		userPosts = append(userPosts, &UserPost)
	}
	err = cursor.Err()
	return
}

func (store *UserPostMongoDBStore) filterNotification(filter interface{}) ([]*domain.Notification, error) {
	cursor, err := store.notificationPostStore.Find(context.TODO(), filter)
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

func (store *UserPostMongoDBStore) InsertNotification(notification *domain.Notification) error {
	_, err := store.notificationPostStore.InsertOne(context.TODO(), notification)
	store.loggerInfo.Logger.Infof("Post_notification_mongodb_store: USCID | UI " + strconv.Itoa(notification.UserId))
	if err != nil {
		println("erorr while inserting notification")
		return err
	}
	//message.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *UserPostMongoDBStore) GetAllUserNotificationsByUserId(id int) ([]*domain.Notification, error) {
	filteringUserNotifications := bson.M{"$or": []bson.M{{"user_id": id}}}
	return store.filterNotification(filteringUserNotifications)
}
