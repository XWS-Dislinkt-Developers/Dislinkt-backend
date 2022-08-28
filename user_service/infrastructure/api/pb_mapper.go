package api

import (
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	"strconv"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       int64(user.ID),
		UserId:   int64(user.UserId),
		Name:     user.Name,
		Username: user.Username,

		Password:         user.Password,
		Email:            user.Email,
		Address:          user.Address,
		Gender:           user.Gender,
		DateOfBirth:      user.DateOfBirth.Format("2006-01-02"),
		PhoneNumber:      user.PhoneNumber,
		IsPrivateProfile: strconv.FormatBool(user.IsPrivateProfile),
		Biography:        user.Biography,

		Skills:    user.Skills,
		Interests: user.Interests,
		Work:      user.Work,
		Education: user.Education,
	}
	return userPb
}
