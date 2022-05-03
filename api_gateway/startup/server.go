package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	cfg "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/startup/config"
	authenticationGw "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	err := authenticationGw.RegisterAuthenticationServiceHandlerFromEndpoint(context.TODO(), server.mux, authenticationEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	// catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	// orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	// shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	// orderingHandler := api.NewOrderingHandler(orderingEmdpoint, catalogueEmdpoint, shippingEmdpoint)
	// orderingHandler.Init(server.mux)
}

// func allowedOrigin(origin string) bool {
// 	if viper.GetString("cors") == "*" {
// 		return true
// 	}
// 	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
// 		return true
// 	}
// 	return false
// }

// func cors(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if allowedOrigin(r.Header.Get("Origin")) {
// 			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
// 			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
// 			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
// 		}
// 		if r.Method == "OPTIONS" {
// 			return
// 		}
// 		h.ServeHTTP(w, r)
// 	})
// }

func (server *Server) Start() {

	// srv := &http.Server{
	// 	Addr:    ":8000",
	// 	Handler: cors(server.mux),
	// }
	//log.Fatal(srv.ListenAndServe())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
