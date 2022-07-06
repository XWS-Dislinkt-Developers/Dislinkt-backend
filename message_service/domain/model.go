package domain

import "github.com/golang-jwt/jwt"

type Message struct {
	Id         int    `bson:"id"`
	SenderId   int    `bson:"sender_id"`
	ReceiverId int    `bson:"receiver_id"`
	Content    string `bson:"content"`
}

type JwtClaims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}
