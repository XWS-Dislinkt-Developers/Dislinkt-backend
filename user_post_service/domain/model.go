package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reaction struct {
	UserId   int  `bson:"user_id"`
	Liked    bool `bson:"liked"`
	Disliked bool `bson:"disliked"`
}

type Comment struct {
	UserId    int       `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
	Text      string    `bson:"text"`
}

type UserPost struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    int                `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Text      string             `bson:"text"`
	ImagePath string             `bson:"image_path"`
	Reactions []Reaction         `bson:"reactions"`
	Comments  []Comment          `bson:"comments"`
}
