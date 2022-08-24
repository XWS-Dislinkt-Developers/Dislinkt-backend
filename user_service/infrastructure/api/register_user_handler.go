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
		//products := mapUpdateProducts(command)
		//err := handler.productService.UpdateQuantityForAll(products)
		//if err != nil {
		//	reply.Type = events.InventoryNotUpdated
		//	break
		//}
		reply.Type = events.ApproveUser
	default:
		reply.Type = events.UnknownCommand
	}

	if reply.Type != events.UnknownCommand {
		_ = handler.replyPublisher.Publish(reply)
	}
}
