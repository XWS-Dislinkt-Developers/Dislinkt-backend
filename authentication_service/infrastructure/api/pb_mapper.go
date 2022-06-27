package api

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       int64(user.ID),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	return userPb
}
