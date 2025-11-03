package authservices

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	userservices "github.com/lunarisnia/crud-example/internal/users/user_services"
)

var secretToken string

type AuthService interface {
	Verify(ctx context.Context, id uint) (string, error)
}

type authServiceImpl struct {
	userService userservices.UserService
}

func NewAuthService(userService userservices.UserService) AuthService {
	return &authServiceImpl{
		userService: userService,
	}
}

func (a authServiceImpl) Verify(ctx context.Context, id uint) (string, error) {
	user, err := a.userService.GetUserById(ctx, id)
	if err != nil {
		return "", err
	}
	if user.ID == 0 {
		return "", errors.New("Unauthorized")
	}

	claim := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	secret, err := FetchToken()
	if err != nil {
		return "", err
	}
	return token.SignedString([]byte(secret))
}

func FetchToken() (string, error) {
	if secretToken == "" {
		token := os.Getenv("JWT_SECRET")
		if token == "" {
			return "", errors.New("Invalid Secret")
		}
		secretToken = token
	}

	return secretToken, nil
}
