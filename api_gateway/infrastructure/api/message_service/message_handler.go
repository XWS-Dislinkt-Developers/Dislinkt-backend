package message_service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type MessageHandler struct {
	messClientAddress string
	//userClientAddress string
}

func NewMessageHandler(messClientAddress string) *MessageHandler {
	return &MessageHandler{
		messClientAddress: messClientAddress,
	}
}

func (handler *MessageHandler) Init(mux *runtime.ServeMux) {
	//err := mux.HandlePath("GET", "/confirmUserAccount/{token}", handler.HandleActivateAccount)
	//if err != nil {
	//	panic(err)
	//}
}

