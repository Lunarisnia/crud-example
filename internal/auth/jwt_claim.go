package auth

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	jwt.RegisteredClaims
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
