package api

import (
	"context"
	"encoding/json"
	"errors"
	pb_message "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/message_service"
	app_message "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"strings"
	"time"
)

type MessageHandler struct {
	pb_message.UnimplementedMessageServiceServer
	message_service      *app_message.MessageService
	notificaiton_service *app_message.NotificationService
	loggerInfo           *logg.Logger
	loggerError          *logg.Logger
}

func NewMessageHandler(message_service *app_message.MessageService, notification_service *app_message.NotificationService, loggerInfo *logg.Logger, loggerError *logg.Logger) *MessageHandler {
	return &MessageHandler{
		message_service:      message_service,
		notificaiton_service: notification_service,
		loggerInfo:           loggerInfo,
		loggerError:          loggerError,
	}
}

// CRUD - READ method(s)
func (handler *MessageHandler) GetAll(ctx context.Context, request *pb_message.GetAllRequest) (*pb_message.MessagesResponse, error) {
	// TODO: This method will be never used?? - Maybe for Admin
	allMessages, err := handler.message_service.GetAll()
	if err != nil {
		handler.loggerError.Logger.Errorf("message_grpc_handler: GetAll - failed method ")
		return nil, err
	}
	response := &pb_message.MessagesResponse{
		Messages: []*pb_message.Message{},
	}
	for _, message := range allMessages {
		current := mapMessage(message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}

func (handler *MessageHandler) GetAllNotifications(ctx context.Context, request *pb_message.GetAllNotificationRequest) (*pb_message.GetAllNotificationResponse, error) {

	header, err := extractHeader(ctx, "authorization")
	if err != nil {
		return &pb_message.GetAllNotificationResponse{}, err
	}
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, err2 := validateToken(token)
	if err2 != nil {
		return &pb_message.GetAllNotificationResponse{}, err2
	}
	notifications, err3 := handler.notificaiton_service.GetAllUserNotificationsByUserId(claims.Id)

	if err3 != nil {
		return &pb_message.GetAllNotificationResponse{}, err3
	}

	response := &pb_message.GetAllNotificationResponse{
		Response: []*pb_message.NotificationResponse{},
	}

	for _, n := range notifications {
		current := &pb_message.NotificationResponse{
			Created: "",
			Content: n.Content,
		}

		response.Response = append(response.Response, current)
	}

	return response, nil
}

func (handler *MessageHandler) GetAllUsersMessagesByUserId(ctx context.Context, request *pb_message.GetAllUsersMessagesRequest) (*pb_message.MessagesResponse, error) {
	// TODO: LoggerInfo? - GetAllSendersMessagesByUserId
	// TODO: authorizationLoggedUser as a separete method?
	header, _ := extractHeader(ctx, "authorization")
	var prefix = "Bearer "
	var token = strings.TrimPrefix(header, prefix)
	claims, _ := validateToken(token)

	userMessages, err := handler.message_service.GetAllUsersMessagesByUserId(claims.Id)
	if err != nil {
		return nil, err
	}
	response := &pb_message.MessagesResponse{
		Messages: []*pb_message.Message{},
	}
	for _, Messages := range userMessages {
		current := mapMessage(Messages)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}

// CRUD - CREATE method(s)
func (handler *MessageHandler) SendMessage(ctx context.Context, request *pb_message.SendMessageRequest) (*pb_message.SendMessageResponse, error) {
	//registerRequestJson, err := decodeBodyToSendMessageRequest(request.Message)
	newMessage := mapNewMessage(request.Message)
	newMessage.SenderId = int(request.Message.SenderId)
	newMessage.ReceiverId = int(request.Message.ReceiverId)
	var new_notif = domain.Notification{UserId: newMessage.ReceiverId, Content: "Dobili ste novu poruku", CreatedAt: timestamppb.Now().AsTime(), Seen: false}
	handler.notificaiton_service.InsertNotification(&new_notif)
	handler.message_service.Insert(newMessage)
	all, _ := handler.notificaiton_service.GetAllUserNotificationsByUserId(newMessage.ReceiverId)
	if all != nil {
		println("nesto")
	}
	return &pb_message.SendMessageResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

// Helper functions - extractHeader, validateToken
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

// Body decoders
func decodeBodyToSendMessageRequest(r io.Reader) (*domain.Message, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var sendMessageRequest domain.Message
	if err := dec.Decode(&sendMessageRequest); err != nil {
		return nil, err
	}
	return &sendMessageRequest, nil
}
