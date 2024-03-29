package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/application"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/infrastructure/api"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/infrastructure/persistence"
	logger "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/logger"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/startup/config"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/intercept"
	authentication "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging/nats"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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
	loggerInfo := logger.InitializeLogger("authentication-service", "INFO")
	loggerError := logger.InitializeLogger("authentication-service", "ERROR")

	postgresClient := server.initPostgresClient()

	userStore := server.initUserStore(postgresClient, loggerInfo, loggerError)

	//sage

	commandPublisher := server.initPublisher("auth.register.command") //server.config.CreateOrderCommandSubject)
	replySubscriber := server.initSubscriber("auth.register.reply", QueueGroup)
	createOrderOrchestrator := server.initRegisterUserOrchestrator(commandPublisher, replySubscriber)

	//kraj sage

	conformationTokenStore := server.initTokenConformationStore(postgresClient, loggerInfo, loggerError)
	passwordRecoveryStore := server.initPasswordRecoveryStore(postgresClient, loggerInfo, loggerError)
	passwordlessLoginStore := server.initPasswordlessLoginStore(postgresClient, loggerInfo, loggerError)

	authService := server.initAuthService(userStore, conformationTokenStore, passwordRecoveryStore, passwordlessLoginStore, loggerInfo, loggerError, createOrderOrchestrator)
	commandSubscriber := server.initSubscriber("auth.register.command", QueueGroup)
	replyPublisher := server.initPublisher("auth.register.reply")
	server.initRegisterUserHandler(authService, replyPublisher, commandSubscriber)

	userHandler := server.initUserHandler(authService, loggerInfo, loggerError)

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
		errInsert, _ := store.Insert(User)
		if errInsert != nil {
			log.Fatal(errInsert)
		}
	}

	return store
}

func (server *Server) initTokenConformationStore(client *gorm.DB, loggerInfo *logger.Logger, loggerError *logger.Logger) domain.ConfirmationTokenStore {
	store, err := persistence.NewConfirmationTokenPostgresStore(client, loggerInfo, loggerError)
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func (server *Server) initPasswordRecoveryStore(client *gorm.DB, loggerInfo *logger.Logger, loggerError *logger.Logger) domain.PasswordRecoveryStore {
	store, err := persistence.NewPasswordRecoveryPostgresStore(client, loggerInfo, loggerError)
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func (server *Server) initPasswordlessLoginStore(client *gorm.DB, loggerInfo *logger.Logger, loggerError *logger.Logger) domain.PasswordlessLoginStore {
	store, err := persistence.NewPasswordlessLoginPostgresStore(client, loggerInfo, loggerError)
	if err != nil {
		log.Fatal(err)
	}
	return store
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
		println("[auth_service][server.go]Publisher   NE    radi")
		println("[auth_service][server.go]Publisher   NE    radi")
		println("[auth_service][server.go]Publisher   NE    radi")
		println("[auth_service][server.go]Publisher   NE    radi")
		println("[auth_service][server.go]Publisher   NE    radi")
	}
	println("[auth_service][server.go]Publisher radi")
	println("[auth_service][server.go]Publisher radi")
	println("[auth_service][server.go]Publisher radi")
	println("[auth_service][server.go]Publisher radi")
	println("[auth_service][server.go]Publisher radi")
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
		println("[auth_service][server.go]Subscriber  NE  radi")
		println("[auth_service][server.go]Subscriber  NE  radi")
		println("[auth_service][server.go]Subscriber  NE  radi")
		println("[auth_service][server.go]Subscriber  NE  radi")
		println("[auth_service][server.go]Subscriber  NE  radi")
	}
	println("[auth_service][server.go]Subscriber radi")
	println("[auth_service][server.go]Subscriber radi")
	println("[auth_service][server.go]Subscriber radi")
	println("[auth_service][server.go]Subscriber radi")
	println("[auth_service][server.go]Subscriber radi")
	return subscriber
}

func (server *Server) initRegisterUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.RegisterUserOrchestrator {
	orchestrator, err := application.NewRegisterUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initAuthService(store domain.UserStore, storeConfToken domain.ConfirmationTokenStore, storePasswordRecovery domain.PasswordRecoveryStore, passwordlessLoginStore domain.PasswordlessLoginStore, loggerInfo *logger.Logger, loggerError *logger.Logger, orchestrator *application.RegisterUserOrchestrator) *application.AuthService {
	return application.NewAuthService(store, storeConfToken, storePasswordRecovery, passwordlessLoginStore, loggerInfo, loggerError, orchestrator)
}

func (server *Server) initRegisterUserHandler(service *application.AuthService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUserHandler(authService *application.AuthService, loggerInfo *logger.Logger, loggerError *logger.Logger) *api.UserHandler {
	return api.NewUserHandler(authService, loggerInfo, loggerError)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				intercept.InterceptToken,
			),
		),
	)
	authentication.RegisterAuthenticationServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
