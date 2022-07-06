package api

import (
	"context"
	"errors"
	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/job_service"
	app_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/job_service/logger"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
	"time"
)

type UserDataHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService

	pb_connection.UnimplementedJobServiceServer
	job_service *app_connection.JobService

	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserDataHandler(job_service *app_connection.JobService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserDataHandler {
	return &UserDataHandler{
		job_service: job_service,
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}

func (handler *UserDataHandler) GetAll(ctx context.Context, request *pb_connection.GetAllRequest) (*pb_connection.GetAllResponse, error) {
	UsersData, err := handler.job_service.GetAllUserData()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: GetAll - failed method ")
		return nil, err
	}

	response := &pb_connection.GetAllResponse{
		UserData: []*pb_connection.UserData{},
	}
	for _, UserConnection := range UsersData {
		current := mapUserData(UserConnection)
		response.UserData = append(response.UserData, current)
	}
	return response, nil
}
func (handler *UserDataHandler) GetDataByUserId(ctx context.Context, id int) (userData domain.UserData) {
	UsersData, _ := handler.job_service.GetUserDataById(id)
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: GAC | UI " + strconv.Itoa(id))
	return *UsersData
}

func (handler *UserDataHandler) GetToken(ctx context.Context, request *pb_connection.GetTokenRequest) (*pb_connection.GetTokenResponse, error) {

	header, err := extractHeader(ctx, "authorization")
	if err != nil {
		return &pb_connection.GetTokenResponse{
			Token: "error: no use token",
		}, nil
	}
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, err := handler.auth_service.ValidateToken(token)
	if err != nil {
		return &pb_connection.GetTokenResponse{
			Token: "error: user token is not valid",
		}, nil
	}

	UsersData, _ := handler.job_service.GetUserDataById(claims.Id)
	var newToken = ""
	if UsersData == nil {
		tempToken, e := handler.job_service.AddToken(int64(claims.Id))
		if e != nil {
			newToken = "error: generating token"
		} else {
			newToken = tempToken
		}
	} else {
		newToken = UsersData.Token
	}
	return &pb_connection.GetTokenResponse{
		Token: newToken,
	}, nil
}

//PostJob
//PostJobCompany

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

func validateToken(signedToken string) (claims *domain.JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&domain.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("Key"), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*domain.JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}
