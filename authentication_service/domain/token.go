package domain

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"time"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/startup/config"
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

var SigningKey = []byte(config.NewConfig().SigningJwtKey)

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwtgo.StandardClaims
}

func keyLookupFunction(token *jwtgo.Token) (interface{}, error) {
	return SigningKey, nil
}

func ParseJwt(tokenStr string) (*jwtgo.Token, *CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenStr, &CustomClaims{}, keyLookupFunction)
	if err != nil {
		return nil, nil, err
	}
	if token == nil {
		return nil, nil, errors.New("Unable to parse token")
	}
	if token.Claims == nil {
		return nil, nil, errors.New("Unable to parse token claims")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		panic("Type Assertion failed")
	}
	return token, claims, err
}
