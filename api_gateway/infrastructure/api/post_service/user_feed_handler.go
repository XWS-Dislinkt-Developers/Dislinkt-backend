package post_service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type UserFeedHandler struct {
	postClientAddress       string
	connectionClientAddress string
}

func NewUserFeedHandler(postClientAddress, connectionClientAddress string) *UserFeedHandler {
	return &UserFeedHandler{
		postClientAddress:       postClientAddress,
		connectionClientAddress: connectionClientAddress,
	}
}

func (handler *UserFeedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/userFeed", handler.HandleUserFeed)
	if err != nil {
		panic(err)
	}
}

func (handler *UserFeedHandler) HandleUserFeed(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//TODO: endpoint koji iz servisa konekcija vraca sve konekcije za ulogovanog korisnika
	//TODO: endpoint koji vraca sve postove iz servisa postova
	//TODO: uzeti sve postove za svaku konekciju, sortirati ih po datumu i vratiti na front
}
