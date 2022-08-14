package startup

import (
	"context"
	"fmt"
	apiAuth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/authentication_service"
	apiConnection "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/connection_service"
	apiPost "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/post_service"
	apiUsers "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/api/user_service"
	cfg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/startup/config"
	authenticationGw "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	userMessageGw "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/message_service"
	userConnectionGw "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_connection_service"
	userPostGw "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_post_service"
	userGw "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {

	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	authenticationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	err1 := authenticationGw.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), server.mux, authenticationEndpoint, opts)
	if err1 != nil {
		panic(err1)
	}

	userPostEndpoint := fmt.Sprintf("%s:%s", server.config.UserPostHost, server.config.UserPostPort)
	err2 := userPostGw.RegisterUserPostServiceHandlerFromEndpoint(context.TODO(), server.mux, userPostEndpoint, opts)
	if err2 != nil {
		panic(err2)
	}

	userConnectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	err3 := userConnectionGw.RegisterUserConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, userConnectionEndpoint, opts)
	if err3 != nil {
		panic(err3)
	}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err4 := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err4 != nil {
		panic(err4)
	}

	userMessageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	err5 := userMessageGw.RegisterMessageServiceHandlerFromEndpoint(context.TODO(), server.mux, userMessageEndpoint, opts)
	if err5 != nil {
		panic(err5)
	}

}

func (server *Server) initCustomHandlers() {
	server.initAccountActivationHandler()
	server.initRegistrationHandler()
	server.initPasswordRecoveryHandler()
	server.initUserFeedHandler()
	server.initUserFeedForPublicProfilesHandler()
	server.initUserConnectionsHandler()
	server.initUserRequestsHandler()
	server.initUserWaitingResponsesHandler()
	server.initUserBlockedHandler()
	server.initUserPostsForProfile()
	server.initSearchForLoggedUserHandler()
}

func (server *Server) initAccountActivationHandler() {
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	accountActivationHandler := apiAuth.NewAccountActivationHandler(authEndpoint, userEndpoint)
	accountActivationHandler.Init(server.mux)
}

func (server *Server) initRegistrationHandler() {
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	userConnectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	registrationHandler := apiAuth.NewRegisterUserHandler(authEndpoint, userEndpoint, userConnectionEndpoint)
	registrationHandler.Init(server.mux)
}

func (server *Server) initPasswordRecoveryHandler() {
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthenticationHost, server.config.AuthenticationPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	passwordReqHandler := apiAuth.NewPasswordRecoveryHandler(authEndpoint, userEndpoint)
	passwordReqHandler.Init(server.mux)
}

func (server *Server) initUserFeedHandler() {
	postEndpoint := fmt.Sprintf("%s:%s", server.config.UserPostHost, server.config.UserPostPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	feedHandler := apiPost.NewUserFeedHandler(postEndpoint, connectionEndpoint)
	feedHandler.Init(server.mux)
}

func (server *Server) initUserFeedForPublicProfilesHandler() {
	postEndpoint := fmt.Sprintf("%s:%s", server.config.UserPostHost, server.config.UserPostPort)
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	feedUsersHandler := apiPost.NewUserFeedForPublicProfilesHandler(postEndpoint, userEndpoint)
	feedUsersHandler.Init(server.mux)
}

func (server *Server) initUserConnectionsHandler() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	connectionUsersHandler := apiConnection.NewUserConnectionsHandler(userEndpoint, connectionEndpoint)
	connectionUsersHandler.Init(server.mux)
}

func (server *Server) initUserRequestsHandler() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	requestsUsersHandler := apiConnection.NewUserRequestHandler(userEndpoint, connectionEndpoint)
	requestsUsersHandler.Init(server.mux)
}

func (server *Server) initUserWaitingResponsesHandler() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	connectionUsersHandler := apiConnection.NewUserWaitingResponseHandler(userEndpoint, connectionEndpoint)
	connectionUsersHandler.Init(server.mux)
}

func (server *Server) initUserBlockedHandler() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	requestsUsersHandler := apiConnection.NewUserBlockedHandler(userEndpoint, connectionEndpoint)
	requestsUsersHandler.Init(server.mux)
}

func (server *Server) initUserPostsForProfile() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postEndpoint := fmt.Sprintf("%s:%s", server.config.UserPostHost, server.config.UserPostPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	userPostsProfileHandler := apiPost.NewUserPostsProfileHandler(userEndpoint, postEndpoint, connectionEndpoint)
	userPostsProfileHandler.Init(server.mux)
}

func (server *Server) initSearchForLoggedUserHandler() {
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.UserConnectionHost, server.config.UserConnectionPort)
	searchUsersHandler := apiUsers.NewUserSearchForLoggedUserHandler(userEndpoint, connectionEndpoint)
	searchUsersHandler.Init(server.mux)
}

func (server *Server) Start() {

	ch := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization"}),
	)

	//HTTPS
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", server.config.Port), server.config.HTTPSServerCertificate, server.config.HTTPSServerKey, ch(server.mux)))
	//HTTP
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), ch(server.mux)))
}
