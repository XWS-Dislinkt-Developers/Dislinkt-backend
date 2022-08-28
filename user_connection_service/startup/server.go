package startup

import (
	"fmt"
	posting "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	saga "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/saga/messaging/nats"
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
	userConnectionService := server.initUserConnectionService(userConnectionStore, loggerInfo, loggerError)
	userNotificationService := server.initNotificationService(userConnectionStore, loggerInfo, loggerError)

	println("[USERSERVICE]Start():init user service")
	commandSubscriber := server.initSubscriber("auth.register.command", QueueGroup)
	replyPublisher := server.initPublisher("auth.register.reply")

	println("[USERSERVICE]Start():sub i pob")
	server.initRegisterUserHandler(userConnectionService, replyPublisher, commandSubscriber)

	userConnectionHandler := server.initUserConnectionHandler(userConnectionService, userNotificationService, loggerInfo, loggerError)
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

func (server *Server) initUserConnectionService(store domain.UserConnectionStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.UserConnectionService {
	return application.NewUserConnectionService(store, loggerInfo, loggerError)
}
func (server *Server) initNotificationService(store domain.UserConnectionStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.NotificationService {
	return application.NewNotificationService(store, loggerInfo, loggerError)
}

func (server *Server) initUserConnectionHandler(service *application.UserConnectionService, notification_service *application.NotificationService, loggerInfo *logger.Logger, loggerError *logger.Logger) *api.UserConnectionHandler {
	return api.NewUserConnectionHandler(service, notification_service, loggerInfo, loggerError)
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

func (server *Server) initRegisterUserHandler(service *application.UserConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewRegisterUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	println("[UserConnectionService][server.go]Publisher radi")
	println("[UserConnectionService][server.go]Publisher radi")
	println("[UserConnectionService][server.go]Publisher radi")
	println("[UserConnectionService][server.go]Publisher radi")
	println("[UserConnectionService][server.go]Publisher radi")
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {

		println("[UserConnectionService][server.go]Subscriber NE radi")
		println("[UserConnectionService][server.go]Subscriber NE radi")
		println("[UserConnectionService][server.go]Subscriber NE radi")
		println("[UserConnectionService][server.go]Subscriber NE radi")
		println("[UserConnectionService][server.go]Subscriber NE radi")
		println("[UserConnectionService][server.go]server.config.NatsHost", server.config.NatsHost)
		println("[UserConnectionService][server.go]server.config.NatsPort", server.config.NatsPort)
		println("[UserConnectionService][server.go]server.config.NatsUser", server.config.NatsUser)
		println("[UserConnectionService][server.go]server.config.NatsPass", server.config.NatsPass)
		println(server.config.NatsHost)
		println(server.config.NatsPort)
		println(server.config.NatsUser)
		println(server.config.NatsPass)
		log.Fatal(err)
	}

	println("[UserConnectionService][server.go]Subscriber radi")
	println("[UserConnectionService][server.go]Subscriber radi")
	println("[UserConnectionService][server.go]Subscriber radi")
	println("[UserConnectionService][server.go]Subscriber radi")
	println("[UserConnectionService][server.go]Subscriber radi")

	return subscriber
}
