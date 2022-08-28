package startup

import (
	"fmt"
	"log"
	"net"

	user "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging/nats"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/infrastructure/persistence"
	logger "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/logger"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/startup/config"
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
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	println("[USERSERVICE]Start():pocetak")
	loggerInfo := logger.InitializeLogger("user-service", "INFO")
	loggerError := logger.InitializeLogger("user-service", "ERROR")

	postgresClient := server.initPostgresClient()
	println("[USERSERVICE]Start():init postgress")
	userStore := server.initUserStore(postgresClient, loggerInfo, loggerError)
	userService := server.initUserService(userStore, loggerInfo, loggerError)
	println("[USERSERVICE]Start():init user service")
	commandSubscriber := server.initSubscriber("auth.register.command", QueueGroup)
	replyPublisher := server.initPublisher("auth.register.reply")

	println("[USERSERVICE]Start():sub i pob")
	server.initRegisterUserHandler(userService, replyPublisher, commandSubscriber)
	println("[USERSERVICE]Start():register handle")
	userHandler := server.initUsersHandler(userService, loggerInfo, loggerError)
	println("[USERSERVICE]Start():init user handle")
	println("[USERSERVICE]Start():sad ide start")
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

func (server *Server) initUserStore(client *gorm.DB, loggerInfo *logger.Logger, loggerError *logger.Logger) domain.UserStore {
	store, err := persistence.NewUserPostgresStore(client, loggerInfo, loggerError)
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
	println("[Userservice][server.go]Publisher radi")
	println("[Userservice][server.go]Publisher radi")
	println("[Userservice][server.go]Publisher radi")
	println("[Userservice][server.go]Publisher radi")
	println("[Userservice][server.go]Publisher radi")
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {

		println("[Userservice][server.go]Subscriber NE radi")
		println("[Userservice][server.go]Subscriber NE radi")
		println("[Userservice][server.go]Subscriber NE radi")
		println("[Userservice][server.go]Subscriber NE radi")
		println("[Userservice][server.go]Subscriber NE radi")
		println("[Userservice][server.go]server.config.NatsHost", server.config.NatsHost)
		println("[Userservice][server.go]server.config.NatsPort", server.config.NatsPort)
		println("[Userservice][server.go]server.config.NatsUser", server.config.NatsUser)
		println("[Userservice][server.go]server.config.NatsPass", server.config.NatsPass)
		println(server.config.NatsHost)
		println(server.config.NatsPort)
		println(server.config.NatsUser)
		println(server.config.NatsPass)
		log.Fatal(err)
	}

	println("[Userservice][server.go]Subscriber radi")
	println("[Userservice][server.go]Subscriber radi")
	println("[Userservice][server.go]Subscriber radi")
	println("[Userservice][server.go]Subscriber radi")
	println("[Userservice][server.go]Subscriber radi")

	return subscriber
}

func (server *Server) initUserService(store domain.UserStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.UserService {
	return application.NewUserService(store, loggerInfo, loggerError)
}

func (server *Server) initUsersHandler(service *application.UserService, loggerInfo *logger.Logger, loggerError *logger.Logger) *api.UsersHandler {
	return api.NewUsersHandler(service, loggerInfo, loggerError)
}

func (server *Server) startGrpcServer(usersHandler *api.UsersHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, usersHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) initRegisterUserHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
