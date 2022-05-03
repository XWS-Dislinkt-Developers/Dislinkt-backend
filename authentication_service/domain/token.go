package domain

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Id       string
	Username string
	jwt.StandardClaims
}
