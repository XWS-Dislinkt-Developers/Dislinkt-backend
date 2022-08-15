package api

import (
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/application"
)

type RegisterUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		userService:       userService,
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
	case events.SaveInUserDB:
		user := mapRegisterUser(command)
		err := handler.userService.Create(user)
		if err != nil {
			reply.Type = events.UserNotSavedDB
		}
		reply.Type = events.UserSavedDB

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
