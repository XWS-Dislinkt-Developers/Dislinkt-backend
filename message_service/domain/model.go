package domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id         int                `bson:"id"`
	SenderId   int                `bson:"sender_id"`
	ReceiverId int                `bson:"receiver_id"`
	Content    string             `bson:"content"`
	Time       primitive.DateTime `bson:"date"`
}

type JwtClaims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}
