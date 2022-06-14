package api

import (
	"context"
	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	app_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/application"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strconv"
	"strings"
)

type UserConnectionHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService

	pb_connection.UnimplementedUserConnectionServiceServer
	connection_service *app_connection.UserConnectionService

	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewUserConnectionHandler(connection_service *app_connection.UserConnectionService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserConnectionHandler {
	return &UserConnectionHandler{
		connection_service: connection_service,
		loggerInfo:         loggerInfo,
		loggerError:        loggerError,
	}
}

func (handler *UserConnectionHandler) GetAll(ctx context.Context, request *pb_connection.GetAllRequest) (*pb_connection.GetAllResponse, error) {
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: GetAll - failed method ")
		return nil, err
	}

	response := &pb_connection.GetAllResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) GetConnectionsByUserId(ctx context.Context, id int) (connections []int) {
	UserConnection, _ := handler.connection_service.GetConnectionsById(id)
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: GetConnectionsByUserId - User with id " + strconv.Itoa(id) + " get all connections")
	return UserConnection.Connections
}

func (handler *UserConnectionHandler) Follow(ctx context.Context, request *pb_connection.FollowRequest) (*pb_connection.FollowResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.Follow(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon follow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: Follow - failed method - User with id " + strconv.Itoa(claims.Id) + " failed to follow user " + strconv.Itoa(int(request.IdUser)))
		return nil, err
	}
	handler.loggerError.Logger.Errorf("User_connection_grpc_handler: Follow - User with id " + strconv.Itoa(claims.Id) + " follow user " + strconv.Itoa(int(request.IdUser)))
	response := &pb_connection.FollowResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}

func (handler *UserConnectionHandler) Unfollow(ctx context.Context, request *pb_connection.FollowRequest) (*pb_connection.FollowResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.Unfollow(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: Unfollow - failed method - User with id " + strconv.Itoa(claims.Id) + " failed to unfollow user " + strconv.Itoa(int(request.IdUser)))
		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: Unfollow - User with id " + strconv.Itoa(claims.Id) + " unfollow user " + strconv.Itoa(int(request.IdUser)))
	response := &pb_connection.FollowResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}

func (handler *UserConnectionHandler) AcceptConnectionRequest(ctx context.Context, request *pb_connection.FollowRequest) (*pb_connection.FollowResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.AcceptConnectionRequest(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: AcceptConnectionRequest - failed method - User with id " + strconv.Itoa(claims.Id) + " couldn't accept connection request ")
		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: AcceptConnectionRequest - User with id " + strconv.Itoa(claims.Id) + " accept connection request ")

	response := &pb_connection.FollowResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}

func (handler *UserConnectionHandler) DeclineConnectionRequest(ctx context.Context, request *pb_connection.FollowRequest) (*pb_connection.FollowResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.DeclineConnectionRequest(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: AcceptConnectionRequest - failed method - User with id " + strconv.Itoa(claims.Id) + " couldn't accept connection request ")

		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: AcceptConnectionRequest - User with id " + strconv.Itoa(claims.Id) + " accept connection request from user id " + strconv.Itoa(int(request.IdUser)))

	response := &pb_connection.FollowResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
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
