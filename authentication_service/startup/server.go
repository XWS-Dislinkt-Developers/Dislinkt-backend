package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/infrastructure/persistence"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/startup/config"
	inventory "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging/nats"
	"google.golang.org/grpc"
	"gorm.io/gorm"
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
	QueueGroup = "authentication_service"
)

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	userStore := server.initUserStore(postgresClient)

	userService := server.initUserService(userStore)

	//commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	//replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	//server.initCreateOrderHandler(userService, replyPublisher, commandSubscriber)

	userHandler := server.initUserHandler(userService)

	server.startGrpcServer(userHandler)
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := persistence.GetClient(
		server.config.UserDBHost, server.config.UserDBUser,
		server.config.UserDBPass, server.config.UserDBName,
		server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *gorm.DB) domain.UserStore {
	store, err := persistence.NewUserPostgresStore(client)
	if err != nil {
		log.Fatal(err)
	}
	store.DeleteAll()
	for _, User := range users {
		err := store.Insert(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

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

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

// func (server *Server) initCreateOrderHandler(service *application.ProductService, publisher saga.Publisher, subscriber saga.Subscriber) {
// 	_, err := api.NewCreateOrderCommandHandler(service, publisher, subscriber)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	inventory.RegisterAuthenticationServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
