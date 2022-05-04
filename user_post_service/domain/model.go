package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reaction struct {
	UserId   primitive.ObjectID `bson:"user_id"`
	IsItLike bool               `bson:"is_it_like"`
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
	ImagePath string             `bson:"text"`
	Reactions []Reaction         `bson:"reactions"`
	Comments  []Comment          `bson:"comments"`
}
