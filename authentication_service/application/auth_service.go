package application

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	gomail "gopkg.in/mail.v2"
	"os"
	"strconv"
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
	passwordlessLoginStore domain.PasswordlessLoginStore
	loggerInfo             *logg.Logger
	loggerError            *logg.Logger
}

func NewAuthService(store domain.UserStore, conformationTokenStore domain.ConfirmationTokenStore, passwordRecoveryStore domain.PasswordRecoveryStore, passwordlessLoginStore domain.PasswordlessLoginStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *AuthService {
	return &AuthService{
		store:                  store,
		conformationTokenStore: conformationTokenStore,
		passwordRecoveryStore:  passwordRecoveryStore,
		passwordlessLoginStore: passwordlessLoginStore,
		loggerInfo:             loggerInfo,
		loggerError:            loggerError,
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
		service.loggerError.Logger.Errorf("Auth_service: GenerateToken - failed method - token can't be generated")
		return "", err
	}
	service.loggerInfo.Logger.Infof("Auth_service: GenerateToken - user with id " + strconv.Itoa(user.ID) + " get token and logged in.")
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
		service.loggerError.Logger.Errorf("Auth_service: Validate - failed method - token isn't valid")
		return
	}

	claims, ok := token.Claims.(*domain.JwtClaims)

	if !ok {
		service.loggerError.Logger.Errorf("Auth_service: Validate - failed method - couldn't parse claims")
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		service.loggerError.Logger.Errorf("Auth_service: Validate - failed method - JWT is expired")

		return nil, errors.New("JWT is expired")
	}
	service.loggerInfo.Logger.Infof("Auth_service: Validate - token is valid.")
	return claims, nil
}

func (service *AuthService) SendEmailForUserAuthentication(user *domain.User) {
	m := gomail.NewMessage()
	m.SetHeader("From", "sammilica99@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Confirm your account")
	token, _ := service.GenerateTokenForAccountConfirmation(user)
	var text = "To confirm your account, please click here : https://localhost:8000/confirmAccount/" + token
	m.SetBody("text/plain", text)
	d := gomail.NewDialer("smtp.gmail.com", 587, "sammilica99@gmail.com", "setmkiwpicaxhmti")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		service.loggerError.Logger.Errorf("Auth_service: SendEmailForUserAuthentication - failed method - email isn't sent.")

		fmt.Println(err)
		panic(err)
	}
	service.loggerInfo.Logger.Infof("Auth_service: SendEmailForUserAuthentication - email for user authentication is sent.")
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
	letterBytes := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP123456789"
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (service *AuthService) generateCode() (token string) {
	letterBytes := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP123456789"
	b := make([]byte, 8)
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

	service.loggerInfo.Logger.Infof("Auth_service: UCA | UI " + strconv.Itoa(User.ID))
	return u, err
}

func (service *AuthService) GetByEmail(email string) (*domain.User, error) {
	return service.store.GetByEmail(email)
}

func (service *AuthService) PasswordRecoveryRequest(email string) error {
	User, _ := service.store.GetByEmail(email)
	recoveryPassword := &domain.PasswordRecovery{
		UserId:       User.ID,
		RecoveryCode: service.generateCode(),
		ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(5) +
			time.Second*time.Duration(0)),
	}
	err := service.passwordRecoveryStore.Insert(recoveryPassword)
	service.loggerInfo.Logger.Infof("Auth_service: PSWRR | UI  " + strconv.Itoa(User.ID))
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
	d := gomail.NewDialer("smtp.gmail.com", 587, "sammilica99@gmail.com", "setmkiwpicaxhmti")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {

		service.loggerError.Logger.Error("Auth_service: EINSRC | UI  " + strconv.Itoa(user.ID))

		fmt.Println(err)
		panic(err)
	}
	service.loggerInfo.Logger.Infof("Auth_service: ERCIS  | UI " + strconv.Itoa(user.ID))

}

func (service *AuthService) sendPasswordlessLoginEmail(user *domain.User, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "sammilica99@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Passwordless login")
	var text = "You're code for login is " + code + ".It will be active next 5 minutes."
	m.SetBody("text/plain", text)
	d := gomail.NewDialer("smtp.gmail.com", 587, "sammilica99@gmail.com", "setmkiwpicaxhmti")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		service.loggerError.Logger.Error("Auth_service: EINSPSWL | UI  " + strconv.Itoa(user.ID))

		fmt.Println(err)
		panic(err)
	}
	service.loggerInfo.Logger.Infof("Auth_service: sendPasswordlessLoginEmail - EPSWLIS | UI " + strconv.Itoa(user.ID))

}

func (service *AuthService) PasswordRecovery(code string, password string) string {
	PasswordRecovery, _ := service.passwordRecoveryStore.GetByRecoveryCode(code)

	if !time.Now().Local().Before(PasswordRecovery.ExpiresAt) {
		service.loggerError.Logger.Error("Auth_service: CPSWLRE ")

		return "Code for password recovery expired."
	}

	service.store.UpdatePassword(PasswordRecovery.UserId, service.HashPassword(password))
	service.passwordRecoveryStore.Delete(PasswordRecovery.UserId)

	return ""
}

func (service *AuthService) PasswordlessLogin(code string) (user *domain.User, err string) {
	passwordlessLogin, _ := service.passwordlessLoginStore.GetByCode(code)

	if passwordlessLogin != nil {
		if !time.Now().Local().Before(passwordlessLogin.ExpiresAt) {
			service.passwordlessLoginStore.Delete(passwordlessLogin.UserId)
			return nil, "Code for login expired."
		}

		user, _ = service.store.GetById(passwordlessLogin.UserId)
		service.passwordlessLoginStore.Delete(passwordlessLogin.UserId)
		service.loggerInfo.Logger.Infof("Auth_service: ULI | UI  " + strconv.Itoa(user.ID))
		return user, "You are now logged in!"
	}
	service.loggerError.Logger.Errorf("Auth_service: CLIW | UI  " + strconv.Itoa(user.ID))

	return nil, "Code for login is wrong."
}

func (service *AuthService) PasswordlessLoginRequest(User *domain.User) error {
	passwordlessLogin := &domain.PasswordlessLogin{
		UserId: User.ID,
		Code:   service.generateCode(),
		ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(5) +
			time.Second*time.Duration(0)),
	}
	err := service.passwordlessLoginStore.Insert(passwordlessLogin)
	service.loggerInfo.Logger.Infof("Auth_service: PSWLR | UI  " + strconv.Itoa(User.ID))
	service.sendPasswordlessLoginEmail(User, passwordlessLogin.Code)
	return err
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
