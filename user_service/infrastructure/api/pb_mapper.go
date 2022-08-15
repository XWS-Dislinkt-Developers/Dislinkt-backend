package api

import (
	pb "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	events "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/register_user"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:       int64(user.ID),
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Address:  user.Address,
		Gender:   user.Gender,
	}
	return userPb
}

func mapRegisterUser(command *events.RegisterUserCommand) *domain.User {
	userPb := &domain.User{
		UserId:   int(int64(command.User.ID)),
		Name:     command.User.Name,
		Username: command.User.Username,
		Password: command.User.Password,
		Email:    command.User.Email,
		Address:  "",
		Gender:   "",
	}
	return userPb
}
func mapRollbackUser(command *events.RegisterUserCommand) *domain.User {
	userPb := &domain.User{
		ID:            int(int64(command.User.ID)),
		Name:          command.User.Name,
		Username:      command.User.Username,
		Password:      command.User.Password,
		Email:         command.User.Email,
		IsItConfirmed: command.User.IsItConfirmed,
		Role:          command.User.Role,
	}
	return userPb
}
