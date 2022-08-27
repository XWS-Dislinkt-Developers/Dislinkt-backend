package api

import (
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
)

type RegisterUserCommandHandler struct {
	userService       *application.UserConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewRegisterUserCommandHandler(productService *application.UserConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) (*RegisterUserCommandHandler, error) {
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
	case events.RegisterUserConnection:
		//println("[USER CONNECTION SERVICE]User servis dobio poruku od orkestratora da registruje korisnika")
		//println("[USER CONNECTION SERVICE]Fali kod za dodavanje korisnika")
		//println("[USER CONNECTION SERVICE]Fali kod za obavestavanje da je dodat korisnik")
		//println("[USER CONNECTION SERVICE]TODO: Dodaj kod za pamcenje korisnika")
		err := handler.userService.RegisterUserConnection(&domain.UserConnection{
			UserId:          command.User.Id,
			Private:         false,
			Connections:     nil,
			Requests:        nil,
			WaitingResponse: nil,
			Blocked:         nil,
		})
		println("[USER CONNECTION SERVICE]Kod za pamcenje korisnika radi")
		if err != nil {
			reply.Type = events.UserConnectionNOTSaved
			println("[USER CONNECTION SERVICE]Postoji greska pri dodavanju u kod za pamcenje konekcije")
			println(err.Error())
		} else {
			println("[USER CONNECTION SERVICE]Sve ok pri dodavanju u kod za konekcije korisnika, UserCONNECTIONSaved poslato")
			//reply.Type = events.UserConnectionNOTSaved
			reply.Type = events.UserCONNECTIONSaved
		}
		//products := mapUpdateProducts(command)
		//err := handler.productService.UpdateQuantityForAll(products)
		//if err != nil {
		//	reply.Type = events.InventoryNotUpdated
		//	break
		//}
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		println("[USER CONNECTION SERVICE]Poslat kod za UserServiceUserSaved or NOT_SAVED")
		_ = handler.replyPublisher.Publish(reply)
	}
}
