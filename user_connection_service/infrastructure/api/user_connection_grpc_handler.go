package api

import (
	"context"
	"errors"
	app_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	pb_auth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pb_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	app_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserConnectionHandler struct {
	pb_auth.UnimplementedAuthenticationServiceServer
	auth_service *app_auth.AuthService
	pb_connection.UnimplementedUserConnectionServiceServer
	connection_service   *app_connection.UserConnectionService
	notificaiton_service *app_connection.NotificationService
	loggerInfo           *logg.Logger
	loggerError          *logg.Logger
}

func NewUserConnectionHandler(connection_service *app_connection.UserConnectionService, notification_service *app_connection.NotificationService, loggerInfo *logg.Logger, loggerError *logg.Logger) *UserConnectionHandler {
	return &UserConnectionHandler{
		connection_service:   connection_service,
		notificaiton_service: notification_service,
		loggerInfo:           loggerInfo,
		loggerError:          loggerError,
	}
}

func (handler *UserConnectionHandler) GetById(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.UserConnection, error) {
	UserConnection, err := handler.connection_service.GetConnectionsById(int(request.IdUser))
	if err != nil {
		return nil, err
	}
	response := mapUserConnection(UserConnection)
	return response, nil
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
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: GAC | UI " + strconv.Itoa(id))
	return UserConnection.Connections
}
func (handler *UserConnectionHandler) GetConnectionsByUser(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.Connections, error) {

	println("[USETCONNECTION_SERVICE]:TRAZI KONEKCIJE KORISNIKA")
	println("[USETCONNECTION_SERVICE]:za korisnika se traze konekcije, id:")
	temp2 := int(request.IdUser)
	s := strconv.Itoa(temp2)
	println("[USETCONNECTION_SERVICE]:Evo sto je konvertovao")
	println(s)

	println("[USETCONNECTION_SERVICE]:TRAZI KONEKCIJE KORISNIKA", string(int(request.IdUser)))
	UserConnection, _ := handler.connection_service.GetConnectionsById(int(request.IdUser))
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: GAC | UI " + strconv.Itoa(int(request.IdUser)))

	response := &pb_connection.Connections{
		Connections: &pb_connection.Connection{
			Connections: []int64{},
		},
	}

	for _, c := range UserConnection.Connections {
		response.Connections.Connections = append(response.Connections.Connections, int64(c))
	}

	return response, nil
}

func (handler *UserConnectionHandler) RegisterUserConnection(ctx context.Context, request *pb_connection.RegisterRequest) (*pb_connection.RegisterResponse, error) {
	var userConnection domain.UserConnection
	userConnection.UserId = int(request.IdUser)
	userConnection.Private = request.IsItPrivate
	handler.connection_service.RegisterUserConnection(&userConnection)

	return &pb_connection.RegisterResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (handler *UserConnectionHandler) Follow(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.Follow(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon follow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTFU | UI " + strconv.Itoa(claims.Id))
		return nil, err
	}
	handler.loggerError.Logger.Errorf("User_connection_grpc_handler: SFU | UI " + strconv.Itoa(claims.Id))
	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) Unfollow(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.Unfollow(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTUFU | UI " + strconv.Itoa(claims.Id))
		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: SUFU | UI " + strconv.Itoa(claims.Id))
	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) AcceptConnectionRequest(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.AcceptConnectionRequest(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTACCR | UI " + strconv.Itoa(claims.Id))
		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: SACCR | UI " + strconv.Itoa(claims.Id))

	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) DeclineConnectionRequest(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.DeclineConnectionRequest(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTDCR | UI " + strconv.Itoa(claims.Id))

		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: SDCR | UI " + strconv.Itoa(claims.Id))

	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) CancelConnectionRequest(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.CancelConnectionRequest(claims.Id, int(request.IdUser))

	//Ovo je trenutno da nam vrati sve iz baze nakon unfollow-a, da bismo lakse ispratili na postmanu, posle nam ne treba
	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTDCR | UI " + strconv.Itoa(claims.Id))

		return nil, err
	}
	handler.loggerInfo.Logger.Infof("User_connection_grpc_handler: SDCR | UI " + strconv.Itoa(claims.Id))

	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) BlockUser(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.BlockUser(claims.Id, int(request.IdUser))

	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTFU | UI " + strconv.Itoa(claims.Id))
		return nil, err
	}
	handler.loggerError.Logger.Errorf("User_connection_grpc_handler: SFU | UI " + strconv.Itoa(claims.Id))
	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}
func (handler *UserConnectionHandler) UnblockUser(ctx context.Context, request *pb_connection.UserIdRequest) (*pb_connection.ConnectionsResponse, error) {

	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := handler.auth_service.ValidateToken(token)

	handler.connection_service.UnblockUser(claims.Id, int(request.IdUser))

	UserConnections, err := handler.connection_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("User_connection_grpc_handler: FTFU | UI " + strconv.Itoa(claims.Id))
		return nil, err
	}
	handler.loggerError.Logger.Errorf("User_connection_grpc_handler: SFU | UI " + strconv.Itoa(claims.Id))
	response := &pb_connection.ConnectionsResponse{
		UserConnections: []*pb_connection.UserConnection{},
	}
	for _, UserConnection := range UserConnections {
		current := mapUserConnection(UserConnection)
		response.UserConnections = append(response.UserConnections, current)
	}
	return response, nil
}

func (handler *UserConnectionHandler) GetAllNotifications(ctx context.Context, request *pb_connection.GetAllNotificationRequest) (*pb_connection.GetAllNotificationResponse, error) {

	header, err := extractHeader(ctx, "authorization")
	if err != nil {
		return &pb_connection.GetAllNotificationResponse{}, err
	}
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, err2 := validateToken(token)
	if err2 != nil {
		return &pb_connection.GetAllNotificationResponse{}, err2
	}
	notifications, err3 := handler.notificaiton_service.GetAllUserNotificationsByUserId(claims.Id)

	if err3 != nil {
		return &pb_connection.GetAllNotificationResponse{}, err3
	}

	response := &pb_connection.GetAllNotificationResponse{
		Response: []*pb_connection.NotificationResponse{},
	}

	for _, n := range notifications {
		current := &pb_connection.NotificationResponse{
			Content:   n.Content,
			CreatedAt: timestamppb.New(n.CreatedAt),
			SenderId:  int64(n.SenderId),
		}

		response.Response = append(response.Response, current)
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
