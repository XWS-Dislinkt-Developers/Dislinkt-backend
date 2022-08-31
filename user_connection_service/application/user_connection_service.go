package application

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

type UserConnectionService struct {
	store       domain.UserConnectionStore
	graphstore  domain.GraphConnectionStore
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserConnectionService(store domain.UserConnectionStore, graphstore domain.GraphConnectionStore, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserConnectionService {
	return &UserConnectionService{
		store:       store,
		graphstore:  graphstore,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (service *UserConnectionService) GetAll() ([]*domain.UserConnection, error) {
	return service.store.GetAll()
}
func (service *UserConnectionService) GetConnectionsById(idUser int) (*domain.UserConnection, error) {
	//return service.store.GetByUserId(idUser)
	s := strconv.Itoa(idUser)
	println("[USETCONNECTION_SERVICE]:TRAZI KONEKCIJE KORISNIKA, u metodi get conn by id: ", s)
	blocke, err := service.graphstore.GetBlockeds(s)

	if blocke == nil {
		println("[USETCONNECTION_SERVICE]:NEMA BLOCKE")

	}

	println("[USETCONNECTION_SERVICE]:TRAZI")
	if err != nil {
		println("[USETCONNECTION_SERVICE]:TRAZI GRESKA1")
		return nil, err
	}

	friends, err := service.graphstore.GetFriends(s)

	if friends == nil {
		println("[USETCONNECTION_SERVICE]:NEMA friends")

	}

	if err != nil {
		println("[USETCONNECTION_SERVICE]:TRAZI GRESKA2")
		return nil, err
	}

	requests, err := service.graphstore.GetFriendRequests(s)
	if err != nil {
		println("[USETCONNECTION_SERVICE]:TRAZI GRESKA3")
		return nil, err
	}

	if requests == nil {
		println("[USETCONNECTION_SERVICE]:NEMA requests")

	}

	waiting, err := service.graphstore.GetWaitingRequests(s)
	if err != nil {

		println("[USETCONNECTION_SERVICE]:TRAZI GRESKA4")
		println(err.Error())
	}

	var tempwaiting []int
	var tempblocked []int
	var tempconnection []int
	var temprequests []int

	isPrivate := service.graphstore.IsUserPrivateDB(s)
	if isPrivate {
		println("[USERCONNECITON_SERVICE]Za profil vraceno da je true")
	} else {
		println("[USERCONNECITON_SERVICE]Za profil vraceno da je false")
	}

	for _, s := range waiting {
		intVar, _ := strconv.Atoi(s.UserID)
		println("[USETCONNECTION_SERVICE]:Iterira waiting")
		tempwaiting = append(tempwaiting, intVar)
	}

	for _, s := range blocke {
		intVar, _ := strconv.Atoi(s.UserID)
		println("[USETCONNECTION_SERVICE]:Iterira 1")
		tempblocked = append(tempblocked, intVar)
	}

	for _, s := range friends {
		intVar, _ := strconv.Atoi(s.UserID)
		println("[USETCONNECTION_SERVICE]:Iterira2")
		tempconnection = append(tempconnection, intVar)
	}

	for _, s := range requests {
		intVar, _ := strconv.Atoi(s.UserID)
		println("[USETCONNECTION_SERVICE]:Iterira3")
		temprequests = append(temprequests, intVar)
	}

	var conn = domain.UserConnection{
		UserId:          idUser,
		Private:         isPrivate,
		Connections:     tempconnection,
		Requests:        temprequests,
		WaitingResponse: tempwaiting, //TODO: what is this?
		Blocked:         tempblocked,
	}

	return &conn, nil
}

func (service *UserConnectionService) RegisterUserConnection(connection *domain.UserConnection) error {
	s := strconv.Itoa(connection.UserId)
	err := service.graphstore.Register(s, true)

	//err := service.store.Insert(connection)
	if err != nil {
		service.loggerError.Logger.Error("User_connection_service: CNSU ")
		println("Error in create method")
		return err
	}
	return nil
}
func (service *UserConnectionService) Follow(idLoggedUser int, idUser int) {
	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	if service.graphstore.IsUserPrivateDB(userToFollows) { //ako je privatan profil onda posalji sahtev
		err := service.graphstore.SendFriendRequest(loggedUser, userToFollows)
		if err != nil {
			println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri slanju zahteva za prijateljstvo")
		} else {
			var temp = domain.Notification{
				UserId:    idUser,
				SenderId:  idLoggedUser,
				Content:   "You have new request for connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}

			service.store.InsertNotification(&temp)
		}
	} else {
		err := service.graphstore.AddFriend(loggedUser, userToFollows)
		if err != nil {
			println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri postajanju prijatelja")
			println(err)
		} else {
			var temp = domain.Notification{
				UserId:    idUser,
				SenderId:  idLoggedUser,
				Content:   "You have new connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}
			var temp2 = domain.Notification{
				UserId:    idLoggedUser,
				SenderId:  idUser,
				Content:   "You have new connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}
			service.store.InsertNotification(&temp)
			service.store.InsertNotification(&temp2)
		}

	}
}

/*
func (service *UserConnectionService) Follow(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	if service.connectionDoesntExist(LoggedUserConnection, UserConnection) && service.requestDoesntExist(LoggedUserConnection, UserConnection) && service.BlockedDoesntExist(UserConnection, LoggedUserConnection) {

		if !service.requestDoesntExist(UserConnection, LoggedUserConnection) {
			UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
			LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
			service.store.UpdateConnections(UserConnection, LoggedUserConnection)
			var temp = domain.Notification{
				UserId:    idUser,
				SenderId:  idLoggedUser,
				Content:   "You have new connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}
			var temp2 = domain.Notification{
				UserId:    idLoggedUser,
				SenderId:  idUser,
				Content:   "You have new connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}
			service.store.InsertNotification(&temp)
			service.store.InsertNotification(&temp2)
		} else if UserConnection.Private {
			UserConnection.Requests = append(UserConnection.Requests, idLoggedUser)
			LoggedUserConnection.WaitingResponse = append(LoggedUserConnection.WaitingResponse, idUser)
			service.store.UpdateRequestConnection(UserConnection)
			service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
			var temp = domain.Notification{
				UserId:    idUser,
				SenderId:  idLoggedUser,
				Content:   "You have new request for connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}

			service.store.InsertNotification(&temp)

		} else {
			if !service.waitingResponseDoesntExist(UserConnection, LoggedUserConnection) || !service.waitingResponseDoesntExist(LoggedUserConnection, UserConnection) {
				UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
				LoggedUserConnection.WaitingResponse = findAndDelete(LoggedUserConnection.WaitingResponse, idUser)
				service.store.UpdateWaitingResponseConnection(UserConnection)
				service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
			}
			if !service.requestDoesntExist(UserConnection, LoggedUserConnection) || !service.requestDoesntExist(LoggedUserConnection, UserConnection) {
				UserConnection.Requests = findAndDelete(UserConnection.Requests, idLoggedUser)
				LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
				service.store.UpdateRequestConnection(UserConnection)
				service.store.UpdateRequestConnection(LoggedUserConnection)
			}
			UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
			LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
			service.store.UpdateConnections(UserConnection, LoggedUserConnection)
			var temp = domain.Notification{
				UserId:    idUser,
				SenderId:  idLoggedUser,
				Content:   "You have new connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}
			var temp2 = domain.Notification{
				UserId:    idLoggedUser,
				SenderId:  idUser,
				Content:   "You have new connection.",
				CreatedAt: timestamppb.Now().AsTime(),
				Seen:      false,
			}
			service.store.InsertNotification(&temp)
			service.store.InsertNotification(&temp2)

		}
	}
}

*/
//func (service *UserConnectionService) Unfollow(idLoggedUser int, idUser int) {
//	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
//	UserConnection, _ := service.store.GetByUserId(idUser)
//
//	UserConnection.Connections = findAndDelete(UserConnection.Connections, idLoggedUser)
//	LoggedUserConnection.Connections = findAndDelete(LoggedUserConnection.Connections, idUser)
//	service.store.UpdateConnections(UserConnection, LoggedUserConnection)
//}

func (service *UserConnectionService) Unfollow(idLoggedUser int, idUser int) {
	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	err := service.graphstore.RemoveFriend(loggedUser, userToFollows)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri uklanjanju prijatelja")
		println(err)
	}
}

func (service *UserConnectionService) AcceptConnectionRequest(idLoggedUser int, idUser int) {

	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	err := service.graphstore.AddFriend(loggedUser, userToFollows)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri prihvatanju prijatelja")
		println(err.Error())
		return
	} else {
		var temp = domain.Notification{
			UserId:    idUser,
			SenderId:  idLoggedUser,
			Content:   "You have new connection.",
			CreatedAt: timestamppb.Now().AsTime(),
			Seen:      false,
		}
		var temp2 = domain.Notification{
			UserId:    idLoggedUser,
			SenderId:  idUser,
			Content:   "You have new connection.",
			CreatedAt: timestamppb.Now().AsTime(),
			Seen:      false,
		}
		service.store.InsertNotification(&temp)
		service.store.InsertNotification(&temp2)

	}

}

/*func (service *UserConnectionService) AcceptConnectionRequest(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
	service.store.UpdateRequestConnection(LoggedUserConnection)
	UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
	service.store.UpdateWaitingResponseConnection(UserConnection)

	UserConnection.Connections = append(UserConnection.Connections, idLoggedUser)
	LoggedUserConnection.Connections = append(LoggedUserConnection.Connections, idUser)
	service.store.UpdateConnections(UserConnection, LoggedUserConnection)

	var temp = domain.Notification{
		UserId:    idUser,
		SenderId:  idLoggedUser,
		Content:   "You have new connection.",
		CreatedAt: timestamppb.Now().AsTime(),
		Seen:      false,
	}
	var temp2 = domain.Notification{
		UserId:    idLoggedUser,
		SenderId:  idUser,
		Content:   "You have new connection.",
		CreatedAt: timestamppb.Now().AsTime(),
		Seen:      false,
	}
	service.store.InsertNotification(&temp)
	service.store.InsertNotification(&temp2)

}*/
func (service *UserConnectionService) DeclineConnectionRequest(idLoggedUser int, idUser int) {
	//LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	//UserConnection, _ := service.store.GetByUserId(idUser)
	//
	//LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
	//service.store.UpdateRequestConnection(LoggedUserConnection)
	//UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
	//service.store.UpdateWaitingResponseConnection(UserConnection)

	//TODO: ?

	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	err := service.graphstore.CancelRequestFromSomeone(loggedUser, userToFollows)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri odbijanju prijateljstva")
		println(err.Error())
		return
	}

}
func (service *UserConnectionService) CancelConnectionRequest(idLoggedUser int, idUser int) {

	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	//LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	//UserConnection, _ := service.store.GetByUserId(idUser)

	//LoggedUserConnection.WaitingResponse = findAndDelete(LoggedUserConnection.WaitingResponse, idUser)
	//service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
	//UserConnection.Requests = findAndDelete(UserConnection.Requests, idLoggedUser)
	//service.store.UpdateRequestConnection(UserConnection)

	err := service.graphstore.UnsendFriendRequest(loggedUser, userToFollows)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri otkazivanju zahteva za prijateljstvo")
		println(err.Error())
		return
	}

}

/*func (service *UserConnectionService) BlockUser(idLoggedUser int, idUser int) {
	LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	UserConnection, _ := service.store.GetByUserId(idUser)

	if service.BlockedDoesntExist(LoggedUserConnection, UserConnection) {
		if !service.connectionDoesntExist(LoggedUserConnection, UserConnection) {
			LoggedUserConnection.Connections = findAndDelete(LoggedUserConnection.Connections, idUser)
			UserConnection.Connections = findAndDelete(UserConnection.Connections, idLoggedUser)
			service.store.UpdateConnections(UserConnection, LoggedUserConnection)
		}
		if !service.requestDoesntExistForBlockingUsers(LoggedUserConnection, UserConnection) {
			LoggedUserConnection.Requests = findAndDelete(LoggedUserConnection.Requests, idUser)
			service.store.UpdateRequestConnection(LoggedUserConnection)
			UserConnection.WaitingResponse = findAndDelete(UserConnection.WaitingResponse, idLoggedUser)
			service.store.UpdateWaitingResponseConnection(UserConnection)
		}
		if !service.waitingResponseDoesntExistForBlockingUsers(LoggedUserConnection, UserConnection) {
			LoggedUserConnection.WaitingResponse = findAndDelete(LoggedUserConnection.WaitingResponse, idUser)
			service.store.UpdateWaitingResponseConnection(LoggedUserConnection)
			UserConnection.Requests = findAndDelete(UserConnection.Requests, idLoggedUser)
			service.store.UpdateRequestConnection(UserConnection)
		}
		LoggedUserConnection.Blocked = append(LoggedUserConnection.Blocked, idUser)
		service.store.UpdateBlockedConnection(LoggedUserConnection)
	}
}*/

func (service *UserConnectionService) BlockUser(idLoggedUser int, idUser int) {

	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	err := service.graphstore.AddBlockUser(loggedUser, userToFollows)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri blokiranju ljudi")
		println(err.Error())
	}

}

func (service *UserConnectionService) UnblockUser(idLoggedUser int, idUser int) {
	//LoggedUserConnection, _ := service.store.GetByUserId(idLoggedUser)
	//
	//LoggedUserConnection.Blocked = findAndDelete(LoggedUserConnection.Blocked, idUser)
	//service.store.UpdateBlockedConnection(LoggedUserConnection)

	loggedUser := strconv.Itoa(idLoggedUser)
	userToFollows := strconv.Itoa(idUser)

	err := service.graphstore.UnblockUser(loggedUser, userToFollows)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri odblokiranju ljudi")
		println(err.Error())
	}

}

func (service *UserConnectionService) ChangePrivacy(idLoggedUser int, private bool) error {
	loggedUser := strconv.Itoa(idLoggedUser)

	err := service.graphstore.ChangePrivacy(loggedUser, private)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri menjenju privatnosti")
		println(err.Error())
	}

	return err

}

func (service *UserConnectionService) GetRecommendation(idLoggedUser int) ([]*domain.UserConn, error) {
	userId := strconv.Itoa(idLoggedUser)

	userConnection, err := service.graphstore.GetRecommendation(userId)
	if err != nil {
		println("[USERCONNECTION_SERVICE][ucs.go]:Greska pri menjenju privatnosti")
		println(err.Error())
	}

	return userConnection, err

}

func (service *UserConnectionService) connectionDoesntExist(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range LoggedUserConnection.Connections {
		if c == UserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) requestDoesntExist(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range UserConnection.Requests {
		if c == LoggedUserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) requestDoesntExistForBlockingUsers(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range LoggedUserConnection.Requests {
		if c == UserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) waitingResponseDoesntExist(UserConnection *domain.UserConnection, LoggedUserConnection *domain.UserConnection) bool {
	for _, c := range UserConnection.Requests {
		if c == LoggedUserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) waitingResponseDoesntExistForBlockingUsers(LoggedUserConnection *domain.UserConnection, UserConnection *domain.UserConnection) bool {
	for _, c := range LoggedUserConnection.WaitingResponse {
		if c == UserConnection.UserId {
			return false
		}
	}
	return true
}
func (service *UserConnectionService) BlockedDoesntExist(loggedUser *domain.UserConnection, user *domain.UserConnection) bool {
	for _, c := range loggedUser.Blocked {
		if c == user.UserId {
			return false
		}
	}
	return true
}

func findAndDelete(s []int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}
