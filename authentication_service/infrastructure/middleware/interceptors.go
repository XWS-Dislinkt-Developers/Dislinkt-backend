package middleware

import (
	"context"
	"fmt"
	jwt "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/authentication_service/domain"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strings"
)

var (
	CommonInterceptors grpc.ServerOption
)

func init() {
	logger := log.WithFields(log.Fields{
		"goapi": "server",
	})

	CommonInterceptors = grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(logger),
		AuthInterceptor(),
		grpc_tags.UnaryServerInterceptor(),
	))
}

func AuthInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//ctx.Value()
		// If the request is for the /auth endpoint, then let the
		// request through without checking for auth.
		//if interceptorAuthUriRegex.MatchString(info.FullMethod) {
		//	return handler(ctx, req)
		//}

		tokenStr, err := grpc_auth.AuthFromMD(ctx, "Bearer")

		if err != nil {
			return req, grpc.Errorf(codes.Unauthenticated, err.Error())
		}

		// Parse the JWT token string into a token object
		token, claims, err := jwt.ParseJwt(tokenStr)
		if err != nil || token == nil {
			return req, grpc.Errorf(codes.Unauthenticated, err.Error())
		} else if !token.Valid {
			return req, grpc.Errorf(codes.Unauthenticated, "Invalid Token")
		}

		FullMethod := strings.Split(info.FullMethod, "/")

		role := claims.Role

		fmt.Println(FullMethod)
		fmt.Println(claims.Role)
		fmt.Println(claims.Username)

		isAuthorized, err := enforce(role, "/register", "post")

		if err != nil {
			fmt.Println(err)
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
