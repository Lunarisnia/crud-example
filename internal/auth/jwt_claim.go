package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	jwt.RegisteredClaims
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ParseJWTClaim(c *gin.Context) (JWTClaim, error) {
	var jwtClaim JWTClaim
	var ok bool
	if token, exist := c.Get("TokenClaim"); exist {
		jwtClaim, ok = token.(JWTClaim)
		if !ok {
			return jwtClaim, errors.New("cast to jwt claim failed.")
		}
	}
	return jwtClaim, nil
}
