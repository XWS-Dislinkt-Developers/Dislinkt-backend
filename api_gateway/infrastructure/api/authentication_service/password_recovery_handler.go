package authentication_service

import (
	"context"
	"encoding/json"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/domain"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/infrastructure/services"
	pbAuth "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	pbUser "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"io"
	"net/http"
)

type PasswordRecoveryHandler struct {
	authClientAddress string
	userClientAddress string
}

func NewPasswordRecoveryHandler(authClientAddress, userClientAddress string) *PasswordRecoveryHandler {
	return &PasswordRecoveryHandler{
		authClientAddress: authClientAddress,
		userClientAddress: userClientAddress,
	}
}

func (handler *PasswordRecoveryHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/userPasswordRecovery", handler.HandlePasswordRecovery)
	if err != nil {
		panic(err)
	}
}

func (handler *PasswordRecoveryHandler) HandlePasswordRecovery(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	passwordRecovery, err := decodeBodyForPasswordRecovery(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Unable to decode request body!"))
		return
	}

	authClient := services.NewAuthenticationClient(handler.authClientAddress)

	response, err := authClient.PasswordRecovery(context.TODO(), &pbAuth.ChangePasswordWithCodeRequest{
		ChangePassword: &pbAuth.ChangePassword{
			Code:            passwordRecovery.Code,
			Password:        passwordRecovery.Password,
			ConfirmPassword: passwordRecovery.ConfirmPassword,
		},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if response.Email != "" {
		userClient := services.NewUserClient(handler.userClientAddress)
		_, errUser := userClient.ChangePassword(context.TODO(), &pbUser.ChangePasswordReq{
			ChangePassword: &pbUser.ChangePassword{
				Email:    response.Email,
				Password: passwordRecovery.Password,
			},
		})
		if errUser != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.Error))
}

func decodeBodyForPasswordRecovery(r io.Reader) (*domain.PasswordRecoveryDTO, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var recovery domain.PasswordRecoveryDTO
	if err := dec.Decode(&recovery); err != nil {
		return nil, err
	}
	return &recovery, nil
}
