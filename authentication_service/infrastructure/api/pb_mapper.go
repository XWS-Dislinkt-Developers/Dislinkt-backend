package api

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       user.ID,
		name:     user.Name,
		username: user.Username,
		password: user.Password,
		email:    user.Email,
		address:  user.Address,
		gender:   user.Gender,
	}
	return userPb
}
