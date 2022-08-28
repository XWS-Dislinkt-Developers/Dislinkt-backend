package domain

type GraphConnectionStore interface {
	GetFriends(id string) ([]UserConn, error)
	GetBlockeds(userID string) ([]UserConn, error)
	Register(userID string, isPublic bool) error
	AddFriend(userIDa, userIDb string) error
	AddBlockUser(userIDa, userIDb string) error
	RemoveFriend(userIDa, userIDb string) error
	UnblockUser(userIDa, userIDb string) error
	GetRecommendation(userID string) ([]*UserConn, error)
	Init()
	SendFriendRequest(userIDa, userIDb string) error
	UnsendFriendRequest(userIDa, userIDb string) error
	//GetConnectionDetail(userIDa, userIDb string) (*pb.ConnectionDetail, error)
	GetFriendRequests(userID string) ([]UserConn, error)
	ChangePrivacy(userID string, private bool) error
	//GetMyContacts(ctx context.Context, request *pb.GetMyContactsRequest) (*pb.ContactsResponse, error)
	DeleteUser(userID string) error
	IsUserPrivateDB(userID string) bool
}
