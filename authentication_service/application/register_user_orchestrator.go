package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
)

type RegisterUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

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

func (o *RegisterUserOrchestrator) Start(user *domain.User) error {
	event := &events.RegisterUserCommand{
		Type: events.SaveInAuthDB,
		User: events.UserDetails{
			ID:            user.ID.Hex(),
			Username:      user.Username,
			Password:      user.Password,
			Email:         user.Email,
			IsItConfirmed: user.IsItConfirmed,
			Role:          user.Role,
		},
	}

	return o.commandPublisher.Publish(event)
}

func (o *RegisterUserOrchestrator) handle(reply *events.RegisterUserReply) {
	command := events.RegisterUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *RegisterUserOrchestrator) nextCommandType(replyType events.RegisterUserReplyType) events.RegisterUserCommandType {
	switch replyType {
	case events.UserSavedAuth:
		return events.SendEmail
	case events.UserNotSavedAuth:
		return events.CancelRegistration //za poslednji korak
	case events.EmailFailed:
		return events.RollbackSavingAuth
	case events.EmailSent:
		return events.SaveInUserDB
	case events.UserNotSavedDB:
		return events.RollbackSavingAuth
	case events.UserSavedDB:
		return events.ApproveRegistration //menja status na sacuvano
	case events.UserAuthRollback:
		return events.CancelRegistration //za poslednji korak
	default:
		return events.UnknownCommand
	}
}
