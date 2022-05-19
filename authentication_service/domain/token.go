package domain

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Id       int
	Username string
	Role     string
	jwt.StandardClaims
}

type ConfirmationToken struct {
	UserId            int    `json:"userId" ,gorm:"unique" validate:"required,userId" `
	ConfirmationToken string `json:"confirmationToken" ,gorm:"unique" validate:"required,confirmationToken" `
}
