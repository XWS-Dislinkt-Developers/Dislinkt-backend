package api

import (
	"context"
	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/job_service"
	app_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/application"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/logger"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
)

type UserDataHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService

	pb_connection.UnimplementedUserDataServiceServer
	user_service *app_connection.UserDataService

	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserDataHandler(user_service *app_connection.UserDataService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserDataHandler {
	return &UserDataHandler{
		user_service: user_service,
		loggerInfo:   loggerInfo,
		loggerError:  loggerError,
	}
}

func (handler *UserDataHandler) GetAll(ctx context.Context, request *pb_connection.GetAllRequest) (*pb_connection.GetAllResponse, error) {
	UsersData, err := handler.user_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: GetAll - failed method ")
		return nil, err
	}

	response := &pb_connection.GetAllResponse{
		UsersData: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UsersData {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserDataHandler) GetDataByUserId(ctx context.Context, id int) (userData domain.UserData) {
	UsersData, _ := handler.user_service.GetDataById(id)
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: GAC | UI " + strconv.Itoa(id))
	return UsersData
}

func extractHeader(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no headers in request")
	}

	authHeaders, ok := md[header]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no header in request")
	}

	if len(authHeaders) != 1 {
		return "", status.Error(codes.Unauthenticated, "more than 1 header in request")
	}

	return authHeaders[0], nil
}
