package application

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"os"
	"time"
	"unicode"

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
	passwordRecoveryStore  domain.PasswordRecoveryStore
}

func NewAuthService(store domain.UserStore, conformationTokenStore domain.ConfirmationTokenStore, passwordRecoveryStore domain.PasswordRecoveryStore) *AuthService {
	return &AuthService{
		store:                  store,
		conformationTokenStore: conformationTokenStore,
		passwordRecoveryStore:  passwordRecoveryStore,
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
	u, err := service.store.ConfirmAccount(User.ID)
	service.conformationTokenStore.Delete(conformationToken.UserId)

	return u, err
}

func (service *AuthService) PasswordRecoveryRequest(email string) error {
	User, _ := service.store.GetByEmail(email)
	recoveryPassword := &domain.PasswordRecovery{
		UserId:       User.ID,
		RecoveryCode: service.generateRandomString(),
		ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(2) +
			time.Second*time.Duration(0)),
	}
	err := service.passwordRecoveryStore.Insert(recoveryPassword)
	service.sendRecoveryCodeEmail(User, recoveryPassword.RecoveryCode)

	return err
}

func (service *AuthService) sendRecoveryCodeEmail(user *domain.User, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "sammilica99@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Password recovery")
	var text = "You're code for password recovery is " + code + ".It will be active next 2 hours."
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

func (service *AuthService) PasswordRecovery(code string, password string) string {
	PasswordRecovery, _ := service.passwordRecoveryStore.GetByRecoveryCode(code)

	if !time.Now().Local().Before(PasswordRecovery.ExpiresAt) {
		return "Code for password recovery expired."
	}

	service.store.UpdatePassword(PasswordRecovery.UserId, service.HashPassword(password))
	service.passwordRecoveryStore.Delete(PasswordRecovery.UserId)

	return ""
}

func (service *AuthService) IsPasswordValid(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	inARowCounter := 0
	passRune := []rune(password)
	characterToCompareWith := string(passRune[0:1])
	if len(password) >= 10 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}

		if string(char) == characterToCompareWith {
			inARowCounter++
		} else {
			inARowCounter = 1
		}
		characterToCompareWith = string(char)

		if inARowCounter >= 3 {
			return false
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func (service *AuthService) CheckForCommonPasswords(password string) bool {
	f, err := os.Open("common_password_list.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if password == scanner.Text() {
			return false
		}
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return true
}
