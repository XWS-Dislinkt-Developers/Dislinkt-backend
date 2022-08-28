package application

import (
	"crypto/tls"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	gomail "gopkg.in/mail.v2"
	"strconv"
)

func (service *AuthService) sendRecoveryCodeEmail(user *domain.User, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "dislinkt10@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Password recovery")
	var text = "You're code for password recovery is " + code + ".It will be active next 2 hours."
	m.SetBody("text/plain", text)
	d := gomail.NewDialer("smtp.gmail.com", 587, "dislinkt10@gmail.com", "lmsxngphcfqpoaol")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {

		service.loggerError.Logger.Error("Auth_service: EINSRC | UI  " + strconv.Itoa(user.ID))
		panic(err)
	}
	service.loggerInfo.Logger.Infof("Auth_service: ERCIS  | UI " + strconv.Itoa(user.ID))

}

func (service *AuthService) sendPasswordlessLoginEmail(user *domain.User, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "dislinkt10@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Passwordless login")
	var text = "You're code for login is " + code + ".It will be active next 5 minutes."
	m.SetBody("text/plain", text)
	d := gomail.NewDialer("smtp.gmail.com", 587, "dislinkt10@gmail.com", "lmsxngphcfqpoaol")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		service.loggerError.Logger.Error("Auth_service: EINSPSWL | UI  " + strconv.Itoa(user.ID))
		panic(err)
	}
	service.loggerInfo.Logger.Infof("Auth_service: sendPasswordlessLoginEmail - EPSWLIS | UI " + strconv.Itoa(user.ID))
}

func (service *AuthService) SendEmailForUserAuthentication(user *domain.User) {
	m := gomail.NewMessage()
	m.SetHeader("From", "dislinkt10@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Confirm your account")
	token, _ := service.GenerateTokenForAccountConfirmation(user)
	var text = "To confirm your account, please click here : https://localhost:8000/confirmUserAccount/" + token
	m.SetBody("text/plain", text)
	d := gomail.NewDialer("smtp.gmail.com", 587, "dislinkt10@gmail.com", "lmsxngphcfqpoaol")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		service.loggerError.Logger.Errorf("Auth_service: FSEFA")
		panic(err)
	}
	service.loggerInfo.Logger.Infof("Auth_service: EFAS")
}
