package domain

import "github.com/golang-jwt/jwt"

type UserData struct {
	Id     int    `bson:"id"`
	UserId int    `bson:"user_id"`
	Token  string `bson:"token"`
}

type JobOffer struct {
	UserId          int      `bson:"user_id"`
	Company         string   `bson:"company"`
	Position        string   `bson:"position"`
	Description     string   `bson:"description"`
	ExperienceLevel string   `bson:"experience_level"`
	Requirements    []string `bson:"requirements"`
}

type JwtClaims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}
