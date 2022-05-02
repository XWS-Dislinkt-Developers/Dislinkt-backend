package services

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"fmt"

	authenticationService "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/common/proto/authentication_service"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewAuthenticationClient(address string) authenticationService.AuthenticationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		fmt.Println("Gateway faild to start", "Failed to start")
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return authenticationService.NewAuthenticationServiceClient(conn)
}
