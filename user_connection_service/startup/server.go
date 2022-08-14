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
	userConnectionService := server.initUserConnectionService(userConnectionStore, loggerInfo, loggerError)
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

func (server *Server) initUserConnectionService(store domain.UserConnectionStore, loggerInfo *logger.Logger, loggerError *logger.Logger) *application.UserConnectionService {
	return application.NewUserConnectionService(store, loggerInfo, loggerError)
}

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
