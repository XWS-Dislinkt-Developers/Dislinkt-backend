package middleware

import (
	"context"
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/api_gateway/jwt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strings"
)

var (
	CommonInterceptors grpc.ServerOption
)

func init() {
	CommonInterceptors = grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		AuthInterceptor(),
		grpc_tags.UnaryServerInterceptor(),
	))
}

func AuthInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		tokenStr, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return req, grpc.Errorf(codes.Unauthenticated, err.Error())
		}

		token, claims, err := jwt.ParseJwt(tokenStr)
		if err != nil || token == nil {
			return req, grpc.Errorf(codes.Unauthenticated, err.Error())
		} else if !token.Valid {
			return req, grpc.Errorf(codes.Unauthenticated, "Invalid Token")
		}

		FullMethod := strings.Split(info.FullMethod, "/")

		role := claims.Role
		isAuthorized, err := Enforce(role, FullMethod[1], FullMethod[2])

		if err != nil {
			return req, grpc.Errorf(codes.Internal, "Internal Server error while authorization!")
		}

		if !isAuthorized {
			return req, grpc.Errorf(codes.PermissionDenied, "Forbidden request!")
		}

		// Find the user
		//user, err := config.AppConfig.GetUserById(claims.Id)
		//if err != nil {
		//	return req, grpc.Errorf(codes.Unauthenticated, "Invalid User Id")
		//}

		newCtx := context.TODO()
		return handler(newCtx, req)
	}
}
