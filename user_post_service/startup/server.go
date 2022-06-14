package startup

import (
	"fmt"
	posting "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/infrastructure/persistence"
	logger "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/logger"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_post_service/startup/config"
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
	QueueGroup = "user_post_service"
)

func (server *Server) Start() {
	mongoClient := server.initMongoClient()

	loggerInfo := logger.InitializeLogger("post-service", "INFO")
	loggerError := logger.InitializeLogger("post-service", "ERROR")
	userPostStore := server.initUserPostStore(mongoClient, loggerInfo, loggerError)

	//commandPublisher := server.initPublisher(server.config.CreateOrderCommandSubject)
	//replySubscriber := server.initSubscriber(server.config.CreateOrderReplySubject, QueueGroup)
	//createOrderOrchestrator := server.initCreateOrderOrchestrator(commandPublisher, replySubscriber)

	userPostService := server.initUserPostService(userPostStore, loggerInfo, loggerError)

	//commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	//replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	// server.initCreateOrderHandler(orderService, replyPublisher, commandSubscriber)

	userPostHandler := server.initUserPostHandler(userPostService, loggerInfo, loggerError)

	server.startGrpcServer(userPostHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserPostDBHost, server.config.UserPostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserPostStore(client *mongo.Client, loggerInfo *logger.Logger, loggerError *logger.Logger) domain.UserPostStore {
	store := persistence.NewUserPostMongoDBStore(client, loggerInfo, loggerError)
	store.DeleteAll()
	for _, userPost := range userPosts {
		err := store.Insert(userPost)
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
func (server *Server) initUserPostService(store domain.UserPostStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.UserPostService {
	return application.NewUserPostService(store, loggerInfo, loggerError)
}

/*
func (server *Server) initCreateUserPostHandler(service *application.UserPostService) {
	_, err := api.NewCreateUserPostHandler(service)
	if err != nil {
		log.Fatal(err)
	}
}
*/
func (server *Server) initUserPostHandler(service *application.UserPostService, loggerInfo *logger.Logger, loggerError *logger.Logger) *api.UserPostHandler {
	return api.NewUserPostHandler(service, loggerInfo, loggerError)
}

func (server *Server) startGrpcServer(userPostHandler *api.UserPostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	posting.RegisterUserPostServiceServer(grpcServer, userPostHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
