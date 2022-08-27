package domain

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type UserConnection struct {
	UserId          int   `bson:"user_id"`
	Private         bool  `bson:"private_profile"`
	Connections     []int `bson:"connections"`
	Requests        []int `bson:"requests"`
	WaitingResponse []int `bson:"waiting_response"`
	Blocked         []int `bson:"blocked"`
}

type Notification struct {
	UserId    int       `bson:"user_id"`
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
