package api

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type RegisterUserCommandHandler struct {
	authService       *application.AuthService
	emailService      *application.EmailService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(authService *application.AuthService, emailService *application.EmailService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		authService:       authService,
		replyPublisher:    publisher,
		emailService:      emailService,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *RegisterUserCommandHandler) handle(command *events.RegisterUserCommand) {
	id, err := primitive.ObjectIDFromHex(command.User.ID)
	Id, _ := strconv.Atoi(id.Hex()) //nisam sigurna da ovo radi
	if err != nil {
		return
	}
	user := &domain.User{ID: Id}

	reply := events.RegisterUserReply{User: command.User}

	switch command.Type {
	case events.ApproveRegistration:
		err := handler.authService.Approve(user) //napravi metodu
		if err != nil {
			return
		}
		reply.Type = events.UserRegistered
	case events.SendEmail:
		_, err := handler.emailService.SendEmailForUserAuthentication(user)
		if err != nil {
			reply.Type = events.EmailFailed
		}
		reply.Type = events.EmailSent
	case events.RollbackSavingAuth:
		handler.authService.DeleteUser(user)

		reply.Type = events.UserAuthRollback

	case events.CancelRegistration:
		err := handler.authService.CancelRegistration(user)
		if err != nil {
			return
		}
		reply.Type = events.UserRegistrationCancelled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
