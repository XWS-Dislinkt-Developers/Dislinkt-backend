package api

import (
	"context"
	"errors"
	pb_message "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/message_service"
	app_connection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

type MessageHandler struct {
	pb_message.UnimplementedMessageServiceServer
	message_service *app_connection.MessageService

	loggerInfo  *logg.Logger
	loggerError *logg.Logger
}

func NewMessageHandler(message_service *app_connection.MessageService, loggerInfo *logg.Logger, loggerError *logg.Logger) *MessageHandler {
	return &MessageHandler{
		message_service: message_service,
		loggerInfo:      loggerInfo,
		loggerError:     loggerError,
	}
}

func (handler *MessageHandler) GetAll(ctx context.Context, request *pb_message.GetAllRequest) (*pb_message.GetAllResponse, error) {
	_, err := handler.message_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("message_grpc_handler: GetAll - failed method ")
		return nil, err
	}

	response := &pb_message.GetAllResponse{}
	//for _, UserConnection := range UserConnections {
	//	current := mapUserConnection(UserConnection)
	//	response.UserConnections = append(response.UserConnections, current)
	//}
	return response, nil
}
func (handler *MessageHandler) GetMessagesByUserId(ctx context.Context, id int) (connections []int) {
	//UserConnection, _ := handler.message_service.GetConnectionsById(id)

	return nil
}

func (handler *MessageHandler) SendMessage(ctx context.Context, request *pb_message.NewMessageRequest) (*pb_message.NewMessageResponse, error) {

	return nil, nil
}

func (handler *MessageHandler) GetMessageWithUser(ctx context.Context, request *pb_message.GetMessageRequest) (*pb_message.GetMessageResponse, error) {

	return nil, nil
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
