package application

import (
	"errors"
	"time"

	domain "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (service *AuthService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes)
}

func (service *AuthService) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *AuthService) GenerateToken(user *domain.User) (signedToken string, err error) {
	claims := &domain.JwtClaims{
		Id:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(8)).Unix(),
			//Issuer:    service.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte("Key"))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (service *AuthService) ValidateToken(signedToken string) (claims *domain.JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&domain.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("Key"), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*domain.JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil

}

func (service *AuthService) CheckIfAdmin(role string) (retVal bool) {
	if role == "admin" {
		return true
	}
	return false
}

func (service *AuthService) CheckIfUser(role string) (retVal bool) {
	if role == "user" {
		return true
	}
	return false
}
