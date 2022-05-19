package application

import (
	"crypto/tls"
	"errors"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"time"

	domain "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

type AuthService struct {
	SecretKey              string
	Issuer                 string
	ExpirationHours        int64
	store                  domain.UserStore
	conformationTokenStore domain.ConfirmationTokenStore
}

func NewAuthService(store domain.UserStore, conformationTokenStore domain.ConfirmationTokenStore) *AuthService {
	return &AuthService{
		store:                  store,
		conformationTokenStore: conformationTokenStore,
	}
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
		Role:     user.Role,
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

func (service *AuthService) SendEmailForUserAuthentication(user *domain.User) {
	m := gomail.NewMessage()

	m.SetHeader("From", "sammilica99@gmail.com")

	m.SetHeader("To", user.Email)

	m.SetHeader("Subject", "Confirm your account")

	token, _ := service.GenerateTokenForAccountConfirmation(user)

	var text = "To confirm your account, please click here : http://localhost:8000/confirmAccount/" + token

	m.SetBody("text/plain", text)

	d := gomail.NewDialer("smtp.gmail.com", 587, "sammilica99@gmail.com", "yearsandyears")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (service *AuthService) GenerateTokenForAccountConfirmation(user *domain.User) (confToken string, err error) {
	token := &domain.ConfirmationToken{
		UserId:            user.ID,
		ConfirmationToken: service.generateRandomString(),
	}

	service.conformationTokenStore.Insert(token)

	return token.ConfirmationToken, nil
}

func (service *AuthService) generateRandomString() (token string) {
	letterBytes := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (service *AuthService) ConfirmAccount(token string) (*domain.User, error) {
	conformationToken, _ := service.conformationTokenStore.GetByConfirmationToken(token)
	User, _ := service.store.GetById(conformationToken.UserId)
	return service.store.ConfirmAccount(User.ID)
}
