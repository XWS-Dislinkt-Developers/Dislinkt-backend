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
	DATABASE   = "userPost"
	COLLECTION = "userPost"
)

type UserPostMongoDBStore struct {
	userPosts   *mongo.Collection
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserPostMongoDBStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.UserPostStore {
	userPosts := client.Database(DATABASE).Collection(COLLECTION)
	return &UserPostMongoDBStore{
		userPosts:   userPosts,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
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
		store.loggerError.Logger.Errorf("User_post_mongodb_store: Insert - failed method - User with id " + strconv.Itoa(userPost.UserId) + " failed while saving his post in database")
		return err
	}
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: Insert - User with id " + strconv.Itoa(userPost.UserId) + " save his post in database")
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
		store.loggerError.Logger.Errorf("User_post_mongodb_store: UpdateComments - failed method - User with id " + strconv.Itoa(userPost.UserId) + " failed while updating his comment in database")
	}
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: Insert - User with id " + strconv.Itoa(userPost.UserId) + " update his comment in database")
	println("success update ")
}

func (store *UserPostMongoDBStore) AddReaction(userPost *domain.UserPost) {
	_, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"reactions", userPost.Reactions}}}})
	if err != nil {
		println("failed reaction update")
		store.loggerError.Logger.Errorf("User_post_mongodb_store: AddReaction - failed method - User with id " + strconv.Itoa(userPost.UserId) + " failed while adding his reaction in database")
	}
	println("success reaction update ")
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: AddReaction - User with id " + strconv.Itoa(userPost.UserId) + " add reaction in database")
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
		store.loggerError.Logger.Errorf("User_post_mongodb_store: UpdateReactions - failed method - User with id " + strconv.Itoa(userPost.UserId) + "failed while updating reaction in database")

	}
	store.loggerInfo.Logger.Infof("User_post_mongodb_store: UpdateReactions - User with id " + strconv.Itoa(userPost.UserId) + " update reaction in database")
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
