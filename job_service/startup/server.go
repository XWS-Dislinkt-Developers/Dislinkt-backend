package startup

import (
	"fmt"
	posting "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	logg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"

	//saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	//"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging/nats"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/infrastructure/persistence"
	logger "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/logger"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_connection_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	// saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	// "github.com/tamararankovic/microservices_demo/common/saga/messaging/nats"
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
	QueueGroup = "user_connection_service"
)

func (server *Server) Start() {
	loggerInfo := logger.InitializeLogger("connection-service", "INFO")
	loggerError := logger.InitializeLogger("connection-service", "ERROR")

	mongoClient := server.initMongoClient()
	userConnectionStore := server.initUserConnectionStore(mongoClient, loggerInfo, loggerError)
	//commandPublisher := server.initPublisher(server.config.CreateOrderCommandSubject)
	//replySubscriber := server.initSubscriber(server.config.CreateOrderReplySubject, QueueGroup)
	//createOrderOrchestrator := server.initCreateOrderOrchestrator(commandPublisher, replySubscriber)

	userConnectionService := server.initUserConnectionService(userConnectionStore, loggerInfo, loggerError)

	//commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	//replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	// server.initCreateOrderHandler(orderService, replyPublisher, commandSubscriber)

	userConnectionHandler := server.initUserConnectionHandler(userConnectionService, loggerInfo, loggerError)

	server.startGrpcServer(userConnectionHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserConnectionDBHost, server.config.UserConnectionDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserConnectionStore(client *mongo.Client, loggerInfo *logg.Logger, loggerError *logg.Logger) domain.UserConnectionStore {
	store := persistence.NewUserConnectionMongoDBStore(client, loggerInfo, loggerError)
	store.DeleteAll()
	for _, userConnection := range userConnections {
		err := store.Insert(userConnection)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

/*
func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateOrderOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.CreateOrderOrchestrator {
	orchestrator, err := application.NewCreateOrderOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}
*/
func (server *Server) initUserConnectionService(store domain.UserConnectionStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.UserConnectionService {
	return application.NewUserConnectionService(store, loggerInfo, loggerError)
}

/*
func (server *Server) initCreateUserPostHandler(service *application.UserPostService) {
	_, err := api.NewCreateUserPostHandler(service)
	if err != nil {
		log.Fatal(err)
	}
}
*/
func (server *Server) initUserConnectionHandler(service *application.UserConnectionService, loggerInfo *logger.Logger, loggerError *logger.Logger) *api.UserConnectionHandler {
	return api.NewUserConnectionHandler(service, loggerInfo, loggerError)
}

func (server *Server) startGrpcServer(userConnectionHandler *api.UserConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	posting.RegisterUserConnectionServiceServer(grpcServer, userConnectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
