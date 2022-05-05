package domain

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Id       int
	Username string
	jwt.StandardClaims
}
