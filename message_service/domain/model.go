package domain

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Message struct {
	// Id         primitive.ObjectID `bson:"id"`
	SenderId   int       `bson:"sender_id"`
	ReceiverId int       `bson:"receiver_id"`
	Content    string    `bson:"content"`
	CreatedAt  time.Time `bson:"created_at"`
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
