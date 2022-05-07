package persistence

import (
	"context"

	"fmt"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "userPost"
	COLLECTION = "userPost"
)

type UserPostMongoDBStore struct {
	userPosts *mongo.Collection
}

func NewUserPostMongoDBStore(client *mongo.Client) domain.UserPostStore {
	userPosts := client.Database(DATABASE).Collection(COLLECTION)
	return &UserPostMongoDBStore{
		userPosts: userPosts,
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
		return err
	}
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
	/*result, _ := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{
		{
			"$set", bson.D{{"comments", userPost.Comments}},
		},
	}) */

	result, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"comments", userPost.Comments}}}})
	if err != nil {
		println("success update ")
	}
	println("failed")
	println(result)
}

func (store *UserPostMongoDBStore) AddReaction(userPost *domain.UserPost) {

	result, err := store.userPosts.UpdateOne(context.TODO(), bson.M{"_id": userPost.Id}, bson.D{{"$set", bson.D{{"reactions", userPost.Reactions}}}})
	if err != nil {
		println("success reaction update ")
	}
	println("failed reaction update")
	println(result)
}
func (store *UserPostMongoDBStore) UpdateReactions(userReaction *domain.Reaction, userPost *domain.UserPost) {

	result, err := store.userPosts.UpdateOne(
		context.TODO(),
		bson.M{"_id": userPost.Id, "reactions.user_id": userReaction.UserId},
		bson.D{
			{"$set", bson.M{"reactions.0.liked": userReaction.Liked, "reactions.0.disliked": userReaction.Disliked}},
		})
	//0 because that's first structure in array
	println("***************************")
	println("USER REACTION LIKE ")
	println(userReaction.Liked)
	println("USER REACTION LIKE ")
	println(userReaction.Disliked)
	println("***************************")
	fmt.Printf("result matched", result.MatchedCount)
	fmt.Printf("result modified ", result.ModifiedCount)
	if err != nil {
		println("ima eror")
	}
	println("nema eror")

}

/* arrayFilters := bson.A{bson.M{"x.Liked": userReaction.Liked}, bson.M{"y.userReaction": step.Name}}
filter := bson.M{"_id.reactions": bson.M{"user_id": userReaction.UserId}}
result, err := store.userPosts.UpdateOne(context.TODO(), filter, bson.D{{"$set", bson.D{{"reactions", userPost.Reactions}}}})
result, err := store.userPosts.FindOneAndUpdate(context.TODO())
if err != nil {
	println("success reaction update ")
}
println("failed reaction update")
println(result) */

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
