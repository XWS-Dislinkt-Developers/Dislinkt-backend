package domain

import (
	"github.com/golang-jwt/jwt"
	"time"
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

type PasswordRecovery struct {
	UserId       int       `json:"userId" ,gorm:"unique" validate:"required,userId" `
	RecoveryCode string    `json:"recoveryCode" ,gorm:"unique" validate:"required,recoveryCode" `
	ExpiresAt    time.Time `json:"expiresAt" `
}

type PasswordlessLogin struct {
	UserId    int       `json:"userId" ,gorm:"unique" validate:"required,userId" `
	Code      string    `json:"code" ,gorm:"unique" validate:"required,cCode" `
	ExpiresAt time.Time `json:"expiresAt" `
}
