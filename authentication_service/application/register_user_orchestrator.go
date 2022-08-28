package application

import (
	domain "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"

	//events "github.com/tamararankovic/microservices_demo/common/saga/register_user"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	//saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	//"github.com/tamararankovic/microservices_demo/ordering_service/domain"
)

//type CreateOrderOrchestrator struct {
//	commandPublisher saga.Publisher
//	replySubscriber  saga.Subscriber
//}

type RegisterUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

//func NewCreateOrderOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CreateOrderOrchestrator, error) {
//	o := &CreateOrderOrchestrator{
//		commandPublisher: publisher,
//		replySubscriber:  subscriber,
//	}
//	err := o.replySubscriber.Subscribe(o.handle)
//	if err != nil {
//		return nil, err
//	}
//	return o, nil
//}

func NewRegisterUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserOrchestrator, error) {
	o := &RegisterUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *RegisterUserOrchestrator) Start(user *domain.UserRegisterData) error {

	println("[auth_service][orkestrator]pocinje orkestrator")
	println("[auth_service][orkestrator]pocinje orkestrator")
	println("[auth_service][orkestrator]pocinje orkestrator")
	println("[auth_service][orkestrator]pocinje orkestrator")
	println("[auth_service][orkestrator]pocinje orkestrator")

	event := &events.RegisterUserCommand{
		Type: events.RegisterUser,
		User: events.UserDetails{
			Id:            user.ID,
			Name:          user.Name,
			Username:      user.Username,
			Password:      user.Password,
			Email:         user.Email,
			IsItConfirmed: user.IsItConfirmed,
			Role:          user.Role,
			Gender:        user.Gender,
			DateOfBirth:   user.DateOfBirth,
		},
	}

	//event := &events.CreateOrderCommand{
	//	Type: events.UpdateInventory,
	//	Order: events.OrderDetails{
	//		Id:      order.Id.Hex(),
	//		Items:   make([]events.OrderItem, 0),
	//		Address: address,
	//	},
	//}
	//for _, item := range order.Items {
	//	eventItem := events.OrderItem{
	//		Product: events.Product{
	//			Id:    item.Product.Id,
	//			Color: events.Color{Code: item.Product.Color.Code},
	//		},
	//		Quantity: item.Quantity,
	//	}
	//	event.Order.Items = append(event.Order.Items, eventItem)
	//}
	return o.commandPublisher.Publish(event)
}

func (o *RegisterUserOrchestrator) handle(reply *events.RegisterUserReply) {
	command := events.RegisterUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *RegisterUserOrchestrator) nextCommandType(reply events.RegisterUserReplyType) events.RegisterUserCommandType {
	switch reply {
	case events.UserServiceUserSaved:
		println("[Authservice][orkestrator]Primljen kod za UserServiceUserSaved salje kod ApproveUser")
		return events.RegisterUserConnection
	case events.UserServiceUserNotSaved:
		println("[Authservice][orkestrator]Primljen kod za UserServiceUserNOTSaved salje kod RollbackUser")
		return events.RollbackUser
	case events.UserConnectionNOTSaved:
		println("[Authservice][orkestrator]Primljen kod za UserConnectionNOTSaved salje kod RollbackUser")
		return events.RollbackUserConnection
	//case events.InventoryUpdated:
	//	return events.ShipOrder
	//case events.InventoryNotUpdated:
	//	return events.CancelOrder
	//case events.InventoryRolledBack:
	//	return events.CancelOrder
	//case events.OrderShippingScheduled:
	//	return events.ApproveOrder
	//case events.OrderShippingNotScheduled:
	//	return events.RollbackInventory
	default:
		return events.UnknownCommand
	}
}
