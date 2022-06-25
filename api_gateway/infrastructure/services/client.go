package services

import (
	authenticationService "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"
	userService "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewAuthenticationClient(address string) authenticationService.AuthenticationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Authentication service: %v", err)
	}
	return authenticationService.NewAuthenticationServiceClient(conn)
}

func NewUserClient(address string) userService.UserServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to User service: %v", err)
	}
	return userService.NewUserServiceClient(conn)
}
