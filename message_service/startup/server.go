package startup

import (
	"fmt"
	posting "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/message_service"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/infrastructure/api"
	persistence "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/infrastructure/persistence"
	logger "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/logger"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/message_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"log"
	"net" //neki komentar
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "message_service"
)

func (server *Server) Start() {
	loggerInfo := logger.InitializeLogger("message-service", "INFO")
	loggerError := logger.InitializeLogger("message-service", "ERROR")
	mongoClient := server.initMongoClient()
	messageStore := server.initMessageStore(mongoClient, loggerInfo, loggerError)
	messageService := server.initMessageService(messageStore, loggerInfo, loggerError)
	notificationService := server.initNotificationService(messageStore, loggerInfo, loggerError)
	messageHandler := server.initMessageHandler(messageService, notificationService, loggerInfo, loggerError)
	server.startGrpcServer(messageHandler)
}

// INITIALIZE - Client, Store(Repository), Service, Handler(Controller)
func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.MessageDBHost, server.config.MessageDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func (server *Server) initMessageStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.MessageStore {
	store := persistence.NewMessagesMongoDBStore(client, loggerInfo, loggerError)
	store.DeleteAll()
	for _, message := range messages {
		err := store.Insert(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}
func (server *Server) initMessageService(store domain.MessageStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.MessageService {
	return application.NewMessageService(store, loggerInfo, loggerError)
}

func (server *Server) initNotificationService(store domain.MessageStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.NotificationService {
	return application.NewNotificationService(store, loggerInfo, loggerError)
}

func (server *Server) initMessageHandler(service *application.MessageService, notification_service *application.NotificationService, loggerInfo *logger.Logger, loggerError *logger.Logger) *api.MessageHandler {
	return api.NewMessageHandler(service, notification_service, loggerInfo, loggerError)
}

// START GRPC HANDLER
func (server *Server) startGrpcServer(userConnectionHandler *api.MessageHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	posting.RegisterMessageServiceServer(grpcServer, userConnectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
