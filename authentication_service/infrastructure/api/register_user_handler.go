package api

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
)

type RegisterUserCommandHandler struct {
	authService       *application.AuthService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(authService *application.AuthService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
	o := &RegisterUserCommandHandler{
		authService:       authService,
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

	id := command.User.Id
	user := domain.User{ID: id}

	switch command.Type {
	case events.ApproveUser:
		println("user approved", user.ID)
		//err := handler.orderService.Approve(order)
		//if err != nil {
		//	return
		//}
		//reply.Type = events.OrderApproved
	case events.RollbackUser:
		err := handler.authService.DeleteUser(user)
		if err != nil {
			return
		}
	default:

	}

	//if reply.Type != events.UnknownReply {
	//	_ = handler.replyPublisher.Publish(reply)
	//}

	//id, err := primitive.ObjectIDFromHex(command.Order.Id)
	//if err != nil {
	//	return
	//}
	//order := &domain.Order{Id: id}
	//
	//reply := events.CreateOrderReply{Order: command.Order}
	//
	//switch command.Type {
	//case events.ApproveOrder:
	//	err := handler.orderService.Approve(order)
	//	if err != nil {
	//		return
	//	}
	//	reply.Type = events.OrderApproved
	//case events.CancelOrder:
	//	err := handler.orderService.Cancel(order)
	//	if err != nil {
	//		return
	//	}
	//	reply.Type = events.OrderCancelled
	//default:
	//	reply.Type = events.UnknownReply
	//}
	//
	//if reply.Type != events.UnknownReply {
	//	_ = handler.replyPublisher.Publish(reply)
	//}
}
