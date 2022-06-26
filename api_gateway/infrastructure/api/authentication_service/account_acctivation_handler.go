package authentication_service

import (
	"context"
	services "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pbAuth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pbUser "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type AccountActivationHandler struct {
	authClientAddress string
	userClientAddress string
}

func NewAccountActivationHandler(authClientAddress, userClientAddress string) *AccountActivationHandler {
	return &AccountActivationHandler{
		authClientAddress: authClientAddress,
		userClientAddress: userClientAddress,
	}
}

func (handler *AccountActivationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/confirmUserAccount/{token}", handler.HandleActivateAccount)
	if err != nil {
		panic(err)
	}
}

func (handler *AccountActivationHandler) HandleActivateAccount(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	token := pathParams["token"]

	authClient := services.NewAuthenticationClient(handler.authClientAddress)

	response, err := authClient.ConfirmAccount(context.TODO(), &pbAuth.ConfirmAccountRequest{
		Token: token,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userClient := services.NewUserClient(handler.userClientAddress)
	_, errUser := userClient.ConfirmAccount(context.TODO(), &pbUser.ConfirmAccountRequest{
		Email: response.Email,
	})
	if errUser != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
