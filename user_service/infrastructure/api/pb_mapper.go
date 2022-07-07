package api

import (
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       int64(user.ID),
		UserId:   int64(user.UserId),
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Address:  user.Address,
		Gender:   user.Gender,
	}
	return userPb
}
