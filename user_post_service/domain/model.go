package domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	UserId    int       `bson:"user_id"`
	CreatedAt time.Time `bson:"created_at"`
	Text      string    `bson:"text"`
}

type UserPost struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    int                `bson:"user_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Text      string             `bson:"text"`
	ImagePath string             `bson:"image_path"`
	Likes     []int              `bson:"likes"`
	Dislikes  []int              `bson:"dislikes"`
	Comments  []Comment          `bson:"comments"`
}

type Notification struct {
	UserId    int       `bson:"user_id"`
	SenderId  int       `bson:"sender_id"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	Seen      bool      `bson:"seen"`
}

type JwtClaims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}
