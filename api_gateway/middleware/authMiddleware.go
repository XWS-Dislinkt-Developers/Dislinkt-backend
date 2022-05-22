package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}

		fmt.Println(r.URL.Path)
		if r.URL.Path == "/registration" || r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			role := "ANONYMOUS"
			isAuthorized, err := Enforce(role, r.URL.Path, r.Method)
			if !isAuthorized {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized request: " + err.Error()))
				next.ServeHTTP(w, r)
				return
			}

			r.Header.Set("role", role)
			next.ServeHTTP(w, r)
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

		isAuthorized, err := Enforce(role, r.URL.Path, r.Method)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error while authorization: " + err.Error()))
			return
		}

		if !isAuthorized {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized request: " + err.Error()))
			next.ServeHTTP(w, r)
			return
		}

		r.Header.Set("username", username)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}

func Enforce(role string, obj string, act string) (bool, error) {
	m, _ := os.Getwd()
	fmt.Println(m)

	enforcer, err := casbin.NewEnforcer("middleware/rbac_model.conf", "middleware/rbac_policy.csv")
	if err != nil {
		return false, fmt.Errorf("failed to create enforcer: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy: %w", err)
	}
	ok, _ := enforcer.Enforce(role, obj, act)
	return ok, nil
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("123456")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
