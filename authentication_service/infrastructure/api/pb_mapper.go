package api

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Address:  user.Address,
		Gender:   user.Gender,
	}
	return userPb
}
