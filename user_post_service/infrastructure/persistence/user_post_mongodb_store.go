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

// USER POSTS

func (store *UserPostMongoDBStore) Get(id primitive.ObjectID) (*domain.UserPost, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}
func (store *UserPostMongoDBStore) GetAll() ([]*domain.UserPost, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}
func (store *UserPostMongoDBStore) GetPostsByUserId(userId int) ([]*domain.UserPost, error) {
	filter := bson.M{"user_id": userId}
	return store.filter(filter)
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
func (store *UserPostMongoDBStore) UpdateLikes(userPost *domain.UserPost) {
	_, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"likes", userPost.Likes}}}})
	if err != nil {
		store.loggerError.Logger.Errorf("User_post_mongodb_store: FTUCID | UI  " + strconv.Itoa(userPost.UserId))
		println("Failed update likes for a post.")
	} else {
		store.loggerInfo.Logger.Infof("User_post_mongodb_store:: USUCID | UI " + strconv.Itoa(userPost.UserId))
	}
}
func (store *UserPostMongoDBStore) UpdateDislikes(userPost *domain.UserPost) {
	_, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"dislikes", userPost.Dislikes}}}})
	if err != nil {
		store.loggerError.Logger.Errorf("User_post_mongodb_store: FTUCID | UI  " + strconv.Itoa(userPost.UserId))
		println("Failed update dislikes for a post.")
	} else {
		store.loggerInfo.Logger.Infof("User_post_mongodb_store:: USUCID | UI " + strconv.Itoa(userPost.UserId))
	}
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

// NOTIFICATIONS

func (store *UserPostMongoDBStore) GetAllUserNotificationsByUserId(id int) ([]*domain.Notification, error) {
	filteringUserNotifications := bson.M{"$or": []bson.M{{"user_id": id}}}
	return store.filterNotification(filteringUserNotifications)
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
