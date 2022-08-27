package api

import (
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	"time"
)

type RegisterUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(productService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		userService:       productService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *RegisterUserCommandHandler) handle(command *events.RegisterUserCommand) {
	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.RegisterUser: //TODO: add code for inserting user to database
		//println("[USERSERVICE]User servis dobio poruku od orkestratora da registruje korisnika")
		//println("[USERSERVICE]Fali kod za dodavanje korisnika")
		//println("[USERSERVICE]Fali kod za obavestavanje da je dodat korisnik")
		//println("[USERSERVICE]TODO: Dodaj kod za pamcenje korisnika")
		dateString := command.User.DateOfBirth
		date, _ := time.Parse("2006-01-02", dateString)

		err := handler.userService.Create(&domain.User{
			UserId:           command.User.Id,
			Name:             command.User.Name,
			Username:         command.User.Username,
			Password:         command.User.Password,
			Email:            command.User.Email,
			Address:          "",
			Gender:           command.User.Gender,
			DateOfBirth:      date,
			Biography:        "",
			PhoneNumber:      "",
			IsPrivateProfile: false,
			Work:             "",
			Education:        "",
			Interests:        "",
			Skills:           "",
			Role:             command.User.Role,
			IsItConfirmed:    false,
		})
		println("[USERSERVICE]Kod za pamcenje korisnika radi")
		if err != nil {
			reply.Type = events.UserServiceUserNotSaved
			println("[USERSERVICE]Postoji greska pri dodavanju u kod za pamcenje korisnika")
			println(err.Error())
		} else {
			println("[USERSERVICE]Sve ok pri dodavanju u kod za pamcenje korisnika, UserServiceUserSaved poslato")
			reply.Type = events.UserServiceUserSaved
		}
		//products := mapUpdateProducts(command)
		//err := handler.productService.UpdateQuantityForAll(products)
		//if err != nil {
		//	reply.Type = events.InventoryNotUpdated
		//	break
		//}
		//
	case events.RollbackUserConnection:
		handler.userService.DeleteUser(command.User.Id)
		reply.Type = events.UnknownReply
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		println("[USERSERVICE]Poslat kod za UserServiceUserSaved or NOT_SAVED")
		_ = handler.replyPublisher.Publish(reply)
	}
}
