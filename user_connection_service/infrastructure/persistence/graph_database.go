package persistence

import (
	"errors"
	"fmt"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type ConnectionDBStore struct {
	connectionDB *neo4j.Driver
	//LoggingService pbLogg.LoggingServiceClient
	//MessageService pbMessage.MessageServiceClient
	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewConnectionGraphDBStore(client *neo4j.Driver, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.GraphConnectionStore {
	return &ConnectionDBStore{
		connectionDB: client,
		loggerInfo:   loggerInfo,
		loggerError:  loggerError,
	}
}

func (store *ConnectionDBStore) Init() {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		errClear := clearGraphDB(transaction)
		if errClear != nil {
			return nil, errClear
		}
		errInit := initGraphDB(transaction)
		return nil, errInit
	})

	if err != nil {
		store.loggerError.Logger.Errorf("[ConnectionService][graphdatabase]Init: Connection to database failed")
		fmt.Println("Connection Graph Database INIT - Failed", err.Error())
	} else {
		fmt.Println("Connection Graph Database INIT - Successfully")
		store.loggerInfo.Logger.Infof("[ConnectionService][graphdatabase]Init: Connection to database successful")
	}

}

func (store *ConnectionDBStore) GetFriends(userID string) ([]domain.UserConn, error) {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	friends, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (this_user:USER) -[:FRIEND]-> (my_friend:USER) WHERE this_user.userID=$uID RETURN my_friend.userID, my_friend.isPrivate",
			map[string]interface{}{"uID": userID})

		if err != nil {
			return nil, err
		}

		var friends []domain.UserConn
		for result.Next() {
			friends = append(friends, domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
		}
		return friends, nil

	})
	if err != nil {
		return nil, err
	}

	return friends.([]domain.UserConn), nil
}

func (store *ConnectionDBStore) GetBlockeds(userID string) ([]domain.UserConn, error) {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	blockedUsers, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (this_user:USER) -[:BLOCK]-> (my_friend:USER) WHERE this_user.userID=$uID RETURN my_friend.userID, my_friend.isPrivate",
			map[string]interface{}{"uID": userID})

		if err != nil {
			return nil, err
		}

		var friends []domain.UserConn
		for result.Next() {
			friends = append(friends, domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
		}
		return friends, nil

	})
	if err != nil {
		return nil, err
	}

	return blockedUsers.([]domain.UserConn), nil

}

func (store *ConnectionDBStore) GetFriendRequests(userID string) ([]domain.UserConn, error) {
	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	friendsRequest, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (this_user:USER) <-[:REQUEST]- (user_requester:USER) WHERE this_user.userID=$uID RETURN user_requester.userID, user_requester.isPrivate",
			map[string]interface{}{"uID": userID})

		if err != nil {
			return nil, err
		}

		var friendsRequest []domain.UserConn
		for result.Next() {
			friendsRequest = append(friendsRequest, domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
		}
		return friendsRequest, nil

	})
	if err != nil {
		return nil, err
	}

	return friendsRequest.([]domain.UserConn), nil
}

func (store *ConnectionDBStore) Register(userID string, isPrivate bool) error {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userID, transaction) {
			return nil, errors.New("User is already registered")
		}

		_, err := transaction.Run(
			"CREATE (new_user:USER{userID:$userID, isPrivate:$isPrivate})",
			map[string]interface{}{"userID": userID, "isPrivate": isPrivate})

		if err != nil {
			return nil, errors.New("Error while creating node")
		}

		return nil, err
	})

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (store *ConnectionDBStore) GetRequests(userID string) ([]domain.UserConn, error) {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	friends, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (this_user:USER) -[:FRIEND]-> (my_friend:USER) WHERE this_user.userID=$uID RETURN my_friend.userID, my_friend.isPrivate",
			map[string]interface{}{"uID": userID})

		if err != nil {
			return nil, err
		}

		var friends []domain.UserConn
		for result.Next() {
			friends = append(friends, domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
		}
		return friends, nil

	})
	if err != nil {
		return nil, err
	}

	return friends.([]domain.UserConn), nil
}

func (store *ConnectionDBStore) DeleteUser(userID string) error {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(userID, transaction) {
			return nil, errors.New("User does not exist")
		}

		_, err := transaction.Run(
			"MATCH (u:USER) WHERE u.userID=$userID DETACH DELETE u ",
			map[string]interface{}{"userID": userID})

		if err != nil {
			return nil, errors.New("Error with deleting user")
		}

		return nil, err
	})

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (store *ConnectionDBStore) AddFriend(userIDa string, userIDb string) error {
	/*
				Dodavanje novog prijatelja je moguce ako:
		         - userA i userB postoji
				 - userA nije prijatelj sa userB
				 - userA nije blokirao userB
			   	 - userA nije blokiran od strane userB
	*/

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
			if checkIfFriendExist(userIDa, userIDb, transaction) || checkIfFriendExist(userIDb, userIDa, transaction) {

				return nil, errors.New("Users are friend")
			} else {
				if checkIfBlockExist(userIDa, userIDb, transaction) || checkIfBlockExist(userIDb, userIDa, transaction) {

					return nil, errors.New("Users are blocked")
				} else {

					isPrivate, err := isUserPrivate(userIDb, transaction)
					if err != nil {
						fmt.Println(err.Error())
						return nil, err
					}
					if isPrivate {
						// ako je profil privatan, onda uspeva samo ako je ovaj profil vec poslao zahtev pre
						if checkIfFriendRequestExist(userIDb, userIDa, transaction) {
							//ok postoji zahtev, mozemo spajati
							removeFriendRequest(userIDb, userIDa, transaction)
						} else {
							return nil, errors.New("Private profile withouth request")
						}
					}

					dateNow := time.Now().Local().Unix()
					result, err := transaction.Run(
						"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
							"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
							"CREATE (u1)-[r1:FRIEND {date: $dateNow}]->(u2) "+
							"CREATE (u2)-[r2:FRIEND {date: $dateNow}]->(u1) "+
							"RETURN r1.date, r2.date",
						map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb, "dateNow": dateNow})

					if err != nil || result == nil {
						return nil, errors.New("Error while creating friends")
					}

					if checkIfFriendRequestExist(userIDa, userIDb, transaction) {
						removeFriendRequest(userIDa, userIDb, transaction)
					}
					if checkIfFriendRequestExist(userIDb, userIDa, transaction) {
						removeFriendRequest(userIDb, userIDa, transaction)
					}
				}
			}
		} else {
			return nil, errors.New("User does not exist")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil

}

func (store *ConnectionDBStore) AddBlockUser(userIDa string, userIDb string) error {

	/*
			UserA moze da blokira UserB ako:
			 - UserA nije vec blokirao UserB
		     - UserB vec nije blokirao prvi UserA
		  	Uspesno blokiranje rezultuje raskidanjem FRIEND veza izmedju ova dva cvora
	*/

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
			if checkIfBlockExist(userIDa, userIDb, transaction) || checkIfBlockExist(userIDb, userIDa, transaction) {
				return nil, errors.New("Users are blocked")
			} else {
				if checkIfFriendExist(userIDa, userIDb, transaction) {
					removeFriend(userIDa, userIDb, transaction)
				}
				if checkIfFriendExist(userIDb, userIDa, transaction) {
					removeFriend(userIDb, userIDa, transaction)
				}
				if checkIfFriendRequestExist(userIDa, userIDb, transaction) {
					removeFriendRequest(userIDa, userIDb, transaction)
				}
				if checkIfFriendRequestExist(userIDb, userIDa, transaction) {
					removeFriendRequest(userIDb, userIDa, transaction)
				}
				blockUser(userIDa, userIDb, transaction)

				return nil, nil
			}

		} else {
			return nil, errors.New("User does not exist")
		}
	})

	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionDBStore) RemoveFriend(userIDa string, userIDb string) error {

	/*
		UserA mora biti prijatelj sa UserB (ne sme biti blokiran)
		UserA izbacuje prijatelja UserB, cepaju se obe prijateljske veze
	*/

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
			if checkIfFriendExist(userIDa, userIDb, transaction) || checkIfFriendExist(userIDb, userIDa, transaction) {

				removeFriend(userIDa, userIDb, transaction)
				removeFriend(userIDb, userIDa, transaction)

			} else {
				return nil, errors.New("users are not friends")
			}
		} else {
			return nil, errors.New("user does not exist")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionDBStore) UnblockUser(userIDa string, userIDb string) error {
	/*
		UserA moze da unblokira useraB samo ako ga je on blokirao
	*/

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
			println("[UserConnection_service][store][unblock]:Korisnici postoje")
			if checkIfBlockExist(userIDb, userIDa, transaction) {
				//actionResult.Msg = "UserB:" + userIDb + " first block UserA:" + userIDa
				//actionResult.Status = 400 //bad request
				return nil, errors.New("You were first blocked")
			} else {
				if checkIfBlockExist(userIDa, userIDb, transaction) {
					if unblockUser(userIDa, userIDb, transaction) {
						//actionResult.Msg = "successfully user IDa:" + userIDa + " unblock user IDb:" + userIDb
						//actionResult.Status = 200
						return nil, nil
					}
				} else {
					//actionResult.Msg = "UserA:" + userIDa + " and UserB:" + userIDb + " are nod blocked"
					//actionResult.Status = 400 //bad request
					return nil, errors.New("Users are not blocked")
				}
			}
		} else {
			//actionResult.Msg = "user does not exist"
			//actionResult.Status = 400 //bad request
			return nil, errors.New("Users dot exit")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (store *ConnectionDBStore) GetRecommendation(userID string) ([]*domain.UserConn, error) {

	/*
		useru koji salje zahtevm, preporucicemo mu 20 prijatelje njegovih prijatelja
		ali necemo mu preporuciti one koje je on blokirao ili koji su njega blokirali

		takodje dobice jos do 20 preporuka ostlaih usera koji se ne nalaze u prvom skupu

		Metoda GetRecommendation vraca ukupno do 40 disjunktih preporuka
			- do 20 preporuka na osnovu prijatelja
			- do 20 preporuka random

	*/

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	recommendation, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		var recommendation []*domain.UserConn

		friendsOfFriends, err1 := getFriendsOfFriendsButNotBlockedRecommendation(userID, transaction)
		if err1 != nil {
			return recommendation, err1
		}

		for _, recommend := range friendsOfFriends {
			recommendation = append(recommendation, recommend)
		}

		famousRecom, err2 := getFriendRecommendation(userID, transaction)
		if err2 != nil {
			return recommendation, err2
		}

		var addNewRecommend bool = true
		for _, recommend := range famousRecom {
			addNewRecommend = true
			for _, r := range recommendation {
				if recommend.UserID == r.UserID {
					addNewRecommend = false
					break
				}
			}
			if addNewRecommend {
				recommendation = append(recommendation, recommend)
			}
		}

		return recommendation, err1

	})
	if err != nil || recommendation == nil {
		return nil, err
	}

	return recommendation.([]*domain.UserConn), nil
}

func (store *ConnectionDBStore) SendFriendRequest(userIDa string, userIDb string) error {
	/*

	 */

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
			if checkIfFriendExist(userIDa, userIDb, transaction) || checkIfFriendExist(userIDb, userIDa, transaction) {

				return nil, errors.New("users are friend")
			} else {
				if checkIfBlockExist(userIDa, userIDb, transaction) || checkIfBlockExist(userIDb, userIDa, transaction) {

					return nil, errors.New("Block exits")
				} else {
					if checkIfFriendRequestExist(userIDa, userIDb, transaction) || checkIfFriendRequestExist(userIDb, userIDa, transaction) {

					} else {

						isPrivate, err := isUserPrivate(userIDb, transaction)
						if err != nil {
							fmt.Println(err.Error())
							return nil, err
						}

						if isPrivate {
							dateNow := time.Now().Local().Unix()
							result, err := transaction.Run(
								"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
									"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
									"CREATE (u1)-[r1:REQUEST {date: $dateNow}]->(u2) "+
									"RETURN r1.date",
								map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb, "dateNow": dateNow})

							if err != nil || result == nil {
								return nil, err
							}
						} else {
							return nil, errors.New("Private profie")
						}
					}
				}
			}
		} else {
			return nil, errors.New("user do not exit")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionDBStore) UnsendFriendRequest(userIDa string, userIDb string) error {
	/*
		UserA moze da povuce zahtev za prijateljstvo samo ako je poslao
	*/

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
			if checkIfBlockExist(userIDa, userIDb, transaction) || checkIfBlockExist(userIDb, userIDa, transaction) {

				return nil, errors.New("User are blocked")
			} else {
				if checkIfFriendRequestExist(userIDa, userIDb, transaction) {
					removeFriendRequest(userIDa, userIDb, transaction)
					return nil, nil
				} else {
					return nil, errors.New("no requests")
				}
			}
		} else {
			return nil, errors.New("user dont exists")
		}

		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionDBStore) CancelRequestFromSomeone(userIDa string, userIDb string) error {
	/*
		UserA moze da ne prihvati zahtev za prijateljstvo
	*/

	if userIDa == userIDb {
		return errors.New("User id are same")
	}

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		_, err := transaction.Run(
			"MATCH (u1:USER{userID: $uIDa})<-[r:REQUEST]-(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
			map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		return err
	}
	return nil
}

//dodatno za getovanje zahteva koje sam poslao sto su na pendingu
//func removeFriendRequest(userIDa, userIDb string, transaction neo4j.Transaction) bool {
//	result, err := transaction.Run(
//		"MATCH (u1:USER{userID: $uIDa})-[r:REQUEST]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
//		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})
//
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	if result != nil && result.Next() {
//		return true
//	}
//	return false
//}

func (store *ConnectionDBStore) GetWaitingRequests(userID string) ([]domain.UserConn, error) {
	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	friendsRequest, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (this_user:USER) -[:REQUEST]-> (user_requester:USER) WHERE this_user.userID=$uID RETURN user_requester.userID, user_requester.isPrivate",
			map[string]interface{}{"uID": userID})

		if err != nil {
			return nil, err
		}

		var friendsRequest []domain.UserConn
		for result.Next() {
			friendsRequest = append(friendsRequest, domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
		}
		return friendsRequest, nil

	})
	if err != nil {
		return nil, err
	}

	return friendsRequest.([]domain.UserConn), nil
}

//kraj zahteva za pending

//func (store *ConnectionDBStore) GetConnectionDetail(userIDa, userIDb string) (*pb.ConnectionDetail, error) {
//
//	/*
//
//	 */
//	if userIDa == userIDb {
//		return &pb.ConnectionDetail{Error: "userIDa is same as userIDb"}, nil
//	}
//
//	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
//	defer session.Close()
//
//	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
//
//		connectionDetail := &pb.ConnectionDetail{UserIDa: userIDa, UserIDb: userIDb}
//
//		// userIDa is not logged in or irrelevant
//		// used for checking if userIDb account is private
//		if userIDa == "000000000000000000000000" {
//			if !checkIfUserExist(userIDb, transaction) {
//				connectionDetail.Error = "user does not exist"
//				return connectionDetail, nil
//			}
//			isPrivate, err := isUserPrivate(userIDb, transaction)
//			if err != nil {
//				connectionDetail.Error = err.Error()
//				return connectionDetail, err
//			}
//			connectionDetail.IsPrivate = isPrivate
//			connectionDetail.Relation = "NO_RELATION"
//			return connectionDetail, nil
//		}
//
//		if checkIfUserExist(userIDa, transaction) && checkIfUserExist(userIDb, transaction) {
//
//			isPrivate, err := isUserPrivate(userIDb, transaction)
//			if err != nil {
//				connectionDetail.Error = err.Error()
//				return connectionDetail, err
//			}
//
//			connectionDetail.IsPrivate = isPrivate
//
//			if checkIfBlockExist(userIDa, userIDb, transaction) {
//				connectionDetail.Relation = "A_BLOCK_B"
//				return connectionDetail, nil
//			}
//			if checkIfBlockExist(userIDb, userIDa, transaction) {
//				connectionDetail.Relation = "B_BLOCK_A"
//				return connectionDetail, nil
//			}
//
//			if checkIfFriendExist(userIDa, userIDb, transaction) || checkIfFriendExist(userIDb, userIDa, transaction) {
//				connectionDetail.Relation = "CONNECTED"
//				return connectionDetail, nil
//			}
//
//			if checkIfFriendRequestExist(userIDa, userIDb, transaction) {
//				connectionDetail.Relation = "PENDING"
//				return connectionDetail, nil
//			}
//			if checkIfFriendRequestExist(userIDb, userIDa, transaction) {
//				connectionDetail.Relation = "ACCEPT"
//				return connectionDetail, nil
//			}
//
//			connectionDetail.Relation = "NO_RELATION"
//
//		} else {
//			connectionDetail.Error = "user does not exist"
//			return connectionDetail, nil
//		}
//
//		return connectionDetail, nil
//	})
//
//	if result == nil {
//		return &pb.ConnectionDetail{Error: "error"}, err
//	} else {
//		store.logg(context.TODO(), "INFO", "GetConnectionDetail", userIDa, "")
//		return result.(*pb.ConnectionDetail), err
//	}
//}

func (store *ConnectionDBStore) ChangePrivacy(userID string, private bool) error {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(userID, transaction) {
			isPrivate, err := isUserPrivate(userID, transaction)
			if err != nil {

				return nil, err
			}

			if isPrivate != private {
				ok, err := setUserPrivate(userID, private, transaction)
				if err != nil {

					return nil, err
				}
				if !ok {
					return nil, errors.New("error updating privacy")
				} else {

					return nil, nil
				}
			} else {
				return nil, errors.New("already has that privacy")
			}

		} else {
			return nil, errors.New("user does not exist")
		}
	})

	if err != nil {
		return err
	}
	return nil
}

func (store *ConnectionDBStore) IsUserPrivateDB(userID string) bool {

	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	p, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {

		private, err := isUserPrivate(userID, transaction)
		if err != nil {
			return false, err
		}

		return private, nil
	})

	if err != nil {
		println("[USERCONNECITON_SERVICE][GraphStore]Greska pri gledanju privatnosi profila")
		return false

	}
	return p.(bool)
}

//func (store *ConnectionDBStore) GetMyContacts(ctx context.Context, request *pb.GetMyContactsRequest) (*pb.ContactsResponse, error) {
//
//	userID := request.UserID
//
//	session := (*store.connectionDB).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
//	defer session.Close()
//
//	contacts, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
//		result, err := transaction.Run(
//			"MATCH (this_user:USER) -[r:FRIEND]-> (my_friend:USER) WHERE this_user.userID=$uID RETURN my_friend.userID, r.msgID ",
//			map[string]interface{}{"uID": userID})
//
//		if err != nil {
//			return nil, err
//		}
//
//		var contacts []*pb.Contact
//		for result.Next() {
//			contacts = append(contacts, &pb.Contact{UserID: result.Record().Values[0].(string), MsgID: result.Record().Values[1].(string)})
//		}
//		return contacts, nil
//
//	})
//	if err != nil {
//		return nil, err
//	}
//	store.logg(context.TODO(), "INFO", "GetMyContacts", userID, "")
//	contactResponse := &pb.ContactsResponse{Contacts: contacts.([]*pb.Contact)}
//	return contactResponse, nil
//
//}

//TODO: HELP FUNCTIONS

func checkIfUserExist(userID string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (existing_uer:USER) WHERE existing_uer.userID = $userID RETURN existing_uer.userID",
		map[string]interface{}{"userID": userID})

	if result != nil && result.Next() && result.Record().Values[0] == userID {
		return true
	}
	return false
}

func isUserPrivate(userID string, transaction neo4j.Transaction) (bool, error) {
	result, err := transaction.Run(
		"MATCH (existing_uer:USER) WHERE existing_uer.userID = $userID RETURN existing_uer.userID, existing_uer.isPrivate",
		map[string]interface{}{"userID": userID})

	if err != nil {
		return true, err
	}

	if result != nil && result.Next() {
		return result.Record().Values[1].(bool), nil
	}
	return true, err
}

func setUserPrivate(userID string, private bool, transaction neo4j.Transaction) (bool, error) {
	result, err := transaction.Run(
		"MATCH (u:USER) WHERE u.userID=$uID SET u.isPrivate=$private RETURN u.isPrivate ",
		map[string]interface{}{"uID": userID, "private": private})
	if err != nil {
		return false, err
	}
	if result != nil && result.Next() {
		return true, nil
	}
	return false, nil
}

func checkIfFriendExist(userIDa string, userIDb string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"MATCH (u1)-[r:FRIEND]->(u2) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func checkIfFriendRequestExist(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"MATCH (u1)-[r:REQUEST]->(u2) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func checkIfBlockExist(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"MATCH (u1)-[r:BLOCK]->(u2) "+
			"RETURN r.date ",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if result != nil && result.Next() {
		return true
	}
	return false
}

func blockUser(userIDa, userIDb string, transaction neo4j.Transaction) bool {

	dateNow := time.Now().Local().Unix()
	result, err := transaction.Run(
		"MATCH (u1:USER) WHERE u1.userID=$uIDa "+
			"MATCH (u2:USER) WHERE u2.userID=$uIDb "+
			"CREATE (u1)-[r:BLOCK {date: $dateNow}]->(u2) "+
			"RETURN r.date",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb, "dateNow": dateNow, "msgID": "newMsgID"})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func removeFriend(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, err := transaction.Run(
		"MATCH (u1:USER{userID: $uIDa})-[r:FRIEND]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func removeFriendRequest(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, err := transaction.Run(
		"MATCH (u1:USER{userID: $uIDa})-[r:REQUEST]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func unblockUser(userIDa, userIDb string, transaction neo4j.Transaction) bool {
	result, err := transaction.Run(
		"MATCH (u1:USER{userID: $uIDa})-[r:BLOCK]->(u2:USER{userID: $uIDb}) DELETE r RETURN u1.userID",
		map[string]interface{}{"uIDa": userIDa, "uIDb": userIDb})

	if err != nil {
		fmt.Println(err)
		return false
	}
	if result != nil && result.Next() {
		return true
	}
	return false
}

func getFriendsOfFriendsButNotBlockedRecommendation(userID string, transaction neo4j.Transaction) ([]*domain.UserConn, error) {
	result, err := transaction.Run(
		"MATCH (u1:USER)-[:FRIEND]->(u2:USER)<-[:FRIEND]-(u3:USER) "+
			"WHERE u1.userID=$uID AND u3.userID<>$uID "+
			"AND NOT exists((u1)-[:FRIEND]-(u3)) "+
			"AND NOT exists((u1)-[:BLOCK]-(u3)) "+
			"RETURN distinct u3.userID, u3.isPrivate "+
			"LIMIT 20 ",
		map[string]interface{}{"uID": userID})

	if err != nil {
		return nil, err
	}

	var recommendation []*domain.UserConn
	for result.Next() {
		recommendation = append(recommendation, &domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
	}
	return recommendation, nil
}

func getFriendRecommendation(userID string, transaction neo4j.Transaction) ([]*domain.UserConn, error) {
	result, err := transaction.Run(
		"MATCH (u1:USER) "+
			"MATCH (u2:USER)-[r:FRIEND]->(:USER) "+
			"WHERE u1.userID=$uID AND u2.userID<>$uID "+
			"AND NOT exists((u1)-[:FRIEND]-(u2)) "+
			"AND NOT exists((u1)-[:BLOCK]-(u2)) "+
			"RETURN distinct u2.userID, u2.isPrivate, COUNT(r) as num_of_friends "+
			"ORDER BY num_of_friends DESC "+
			"LIMIT 20 ",
		map[string]interface{}{"uID": userID})

	if err != nil {
		return nil, err
	}

	var recommendation []*domain.UserConn
	for result.Next() {
		recommendation = append(recommendation, &domain.UserConn{UserID: result.Record().Values[0].(string), IsPrivate: result.Record().Values[1].(bool)})
	}
	return recommendation, nil
}

func clearGraphDB(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"MATCH (n) DETACH DELETE n",
		map[string]interface{}{})
	return err
}

func initGraphDB(transaction neo4j.Transaction) error {
	_, err := transaction.Run(
		"CREATE  (pera:USER{userID: \"1\", isPrivate : false}),  (joka:USER{userID: \"2\", isPrivate : false}),  (marko:USER{userID: \"3\", isPrivate : false}),(zeksa:USER{userID: \"4\", isPrivate : false}),(sanja:USER{userID: \"5\", isPrivate : false}),(tanja:USER{userID: \"6\", isPrivate : true}),(lale:USER{userID: \"7\", isPrivate : false}),(nena:USER{userID: \"8\", isPrivate : false}),(admin:USER{userID: \"9\", isPrivate : false}),   (pera) -[:FRIEND]-> (marko),  (pera) <-[:FRIEND]- (marko),  (pera) -[:FRIEND]-> (joka),  (pera) <-[:FRIEND]- (joka), (marko) -[:BLOCK]-> (joka),  (joka) -[:BLOCK]-> (marko)  ",
		map[string]interface{}{})

	return err
}
