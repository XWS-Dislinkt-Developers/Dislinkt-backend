package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		username := claims.(jwt.MapClaims)["username"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		if role == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized request: " + err.Error()))
			return
		}

		isAuthorized, err := enforce(role, r.URL.Path, r.Method)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error while authorization: " + err.Error()))
			return
		}

		if !isAuthorized {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Forbidden request: " + err.Error()))
			return
		}

		r.Header.Set("username", username)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}

func enforce(role string, obj string, act string) (bool, error) {
	enforcer, err := casbin.NewEnforcer("infrastructure/middleware/rbac_model.conf", "infrastructure/middleware/rbac_policy.csv")
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, _ := enforcer.Enforce(role, obj, act)
	return ok, nil
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := "123456"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
