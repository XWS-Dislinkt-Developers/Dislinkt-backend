package api

import (
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	"strconv"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		// Cannot be change
		Id:   int64(user.UserId),
		Role: user.Role,
		// Personal data - can be updated
		Name:        user.Name,
		Username:    user.Username,
		Password:    user.Password,
		Email:       user.Email,
		Address:     user.Address,
		Gender:      user.Gender,
		DateOfBirth: user.DateOfBirth.Format("2006-01-02"),
		Biography:   user.Biography,
		// Skills and inters - can be updated
		Skills:    user.Skills,
		Interests: user.Interests,
		// Work and education - can be updated
		Work:      user.Work,
		Education: user.Education,
		// Privacy - can be updated
		IsPrivateProfile: strconv.FormatBool(user.IsPrivateProfile),
	}
	return userPb
}
