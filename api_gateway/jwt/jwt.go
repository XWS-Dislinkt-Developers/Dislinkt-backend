package jwt

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

var SigningKey = []byte("123456")

type CustomClaims struct {
	username string `json:"username"`
	Role     string `json:"role"`
	jwtgo.StandardClaims
}

func CreateJwtWithIdRole(id string, role string, secondsToExpiration int64) (string, error) {
	now := time.Now()
	claims := CustomClaims{
		id,
		role,
		jwtgo.StandardClaims{
			Issuer:    "apiservice",
			Audience:  "apiservice",
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Second * time.Duration(secondsToExpiration)).Unix(),
		},
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	ss, err := token.SignedString(SigningKey)

	return ss, err
}

func keyLookupFunction(token *jwtgo.Token) (interface{}, error) {
	return SigningKey, nil
}

func ParseJwt(tokenStr string) (*jwtgo.Token, *CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenStr, &CustomClaims{}, keyLookupFunction)
	if err != nil {
		return nil, nil, err
	}
	if token == nil {
		return nil, nil, errors.New("Unable to parse token")
	}
	if token.Claims == nil {
		return nil, nil, errors.New("Unable to parse token claims")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		panic("Type Assertion failed")
	}
	return token, claims, err
}
