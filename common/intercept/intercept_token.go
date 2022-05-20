package intercept

import (
	"context"
	"fmt"
	authentication "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

var nonAuthMethods map[string]bool = map[string]bool{
	"/authentication.AuthenticationService/UpdateUserWorkEducation":   true,
	"/authentication.AuthenticationService/UpdateUserSkillsInterests": true,
	"/authentication.AuthenticationService/UpdatePersonalData":        true,
	"/authentication.AuthenticationService/GetAll":                    true,
	"/authentication.AuthenticationService/Register":                  true,
	"/authentication.AuthenticationService/Login":                     true,
	"/authentication.AuthenticationService/FindUser":                  true,
	"/authentication.AuthenticationService/ConfirmAccount":            true,
	"/authentication.AuthenticationService/PasswordRecoveryRequest":   true,
	"/authentication.AuthenticationService/PasswordRecovery":          true,
}

func NewAuthenticationClient(address string) authentication.AuthenticationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		fmt.Println("Gateway faild to start", "Failed to start")
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return authentication.NewAuthenticationServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func InterceptToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	method, _ := grpc.Method(ctx)
	println(method)
	if nonAuthMethods[method] == true {
		return handler(ctx, req)
	}

	auth, err := extractHeader(ctx, "authorization")
	if err != nil {
		return ctx, err
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(auth, prefix) {
		return ctx, status.Error(codes.Unauthenticated, `missing "Bearer " prefix in "authorization" header`)
	}

	authClient := NewAuthenticationClient("localhost:8001")
	var token = strings.TrimPrefix(auth, prefix)
	result, err := authClient.Validate(context.TODO(), &authentication.ValidateRequest{Token: token})
	if result.Status != 200 {
		return ctx, status.Error(codes.Unauthenticated, "invalid token")
	}

	return handler(ctx, req)
}

func extractHeader(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no headers in request")
	}

	authHeaders, ok := md[header]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "no header in request")
	}

	if len(authHeaders) != 1 {
		return "", status.Error(codes.Unauthenticated, "more than 1 header in request")
	}

	return authHeaders[0], nil
}
