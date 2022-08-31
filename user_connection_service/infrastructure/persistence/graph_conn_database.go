package persistence

import (
	"errors"
	"fmt"
	//pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	//pbLogg "github.com/XWS-BSEP-TIM2/dislinkt-backend/common/proto/logging_service"
	//pbMessage "github.com/XWS-BSEP-TIM2/dislinkt-backend/common/proto/message_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConnectionDBStore2 struct {
	connectionDB *neo4j.Driver
	//LoggingService pbLogg.LoggingServiceClient
	//MessageService pbMessage.MessageServiceClient
}

//TODO:METHODS TO IMPLEMENT FOR DATABASE

func (store *ConnectionDBStore) GetByUserId(id int) (*domain.UserConnection, error) {
	//TODO implement me
	panic("implement me")
}

func (store *ConnectionDBStore) GetAll() ([]*domain.UserConnection, error) {
	//TODO implement me
	panic("implement me")
}

func (store *ConnectionDBStore) Insert(userConnection *domain.UserConn) error {
	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		//actionResult := &pb.ActionResult{}

		isPrivate := userConnection.IsPrivate
		userID := userConnection.UserID
		if checkIfUserExist2(userConnection.UserID, transaction) {

			return nil, errors.New("User is registered!")
		}

		_, err := transaction.Run(
			"CREATE (new_user:USER{userID:$userID, isPrivate:$isPrivate})",
			map[string]interface{}{"userID": userID, "isPrivate": isPrivate})

		if err != nil {
			return nil, errors.New("Error while creating new node")
		}

		return nil, nil
	})

	if err != nil {
		println("[COONECTION SERVICE][GRAPHDATABASE]Problem with registering new connection")
		println(err)
		return err
	} else {
		println("[COONECTION SERVICE][GRAPHDATABASE]Registered new connection")

		return nil
	}
}

func (store *ConnectionDBStore) DeleteAll() {
	//TODO implement me
	panic("implement me")
}

func (store *ConnectionDBStore) UpdateRequestConnection(userConnection *domain.UserConnection) {
	//TODO implement me
	panic("implement me")
}

func (store *ConnectionDBStore) UpdateConnections(userConnection *domain.UserConnection, loggedUserConnection *domain.UserConnection) {
	//TODO implement me
	panic("implement me")
}

func (store *ConnectionDBStore) UpdateBlockedConnection(connection *domain.UserConnection) {
	//TODO implement me
	panic("implement me")
}

func (store *ConnectionDBStore) UpdateWaitingResponseConnection(connection *domain.UserConnection) {
	//TODO implement me
	panic("implement me")
}

//func NewConnectionDBStore(client *neo4j.Driver, loggingService pbLogg.LoggingServiceClient, messageService pbMessage.MessageServiceClient) domain.GraphConnectionStore {
//	return &ConnectionDBStore{
//		connectionDB: client,
//		//LoggingService: loggingService,
//		//MessageService: messageService,
//	}
//}

func NewConnectionDBStore2(client *neo4j.Driver) domain.GraphConnectionStore {
	return &ConnectionDBStore{
		connectionDB: client,
		//LoggingService: loggingService,
		//MessageService: messageService,
	}
}

func (store *ConnectionDBStore) Init2() {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		errClear := clearGraphDB2(transaction)
		if errClear != nil {
			return nil, errClear
		}
		errInit := initGraphDB2(transaction)
		return nil, errInit
	})

	if err != nil {
		fmt.Println("Connection Graph Database INIT - Failed", err.Error())
	} else {
		fmt.Println("Connection Graph Database INIT - Successfully")
	}

}

//TODO:METHODS FOR HELP, FROM TRANSACTIONS

func clearGraphDB2(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"MATCH (n) DETACH DELETE n",
		map[string]interface{}{})
	return err
}

func initGraphDB2(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"CREATE  (pera:USER{userID: \"1\", isPrivate : false}),  (marko:USER{userID: \"2\", isPrivate : false}),  (joka:USER{userID: \"3\", isPrivate : true}),   (pera) -[:FRIEND]-> (marko),  (pera) <-[:FRIEND]- (marko),  (pera) -[:FRIEND]-> (joka),  (pera) <-[:FRIEND]- (joka),       (marko) -[:BLOCK]-> (joka),  (joka) -[:BLOCK]-> (marko)  ",
		map[string]interface{}{})
	return err
}

func checkIfUserExist2(userID string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (existing_uer:USER) WHERE existing_uer.userID = $userID RETURN existing_uer.userID",
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}
