package api

import (
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/job_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/domain"
)

func mapUserData(userdata *domain.UserData) *pb_connection.UserData {
	userDataPb := &pb_connection.UserData{
		Id:     int64(userdata.Id),
		UserId: int64(userdata.UserId),
		Token:  userdata.Token,
	}

	return userDataPb
}
