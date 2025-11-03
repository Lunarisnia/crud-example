package authcontrollers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	authdto "github.com/lunarisnia/crud-example/internal/auth/auth_dto"
	authservices "github.com/lunarisnia/crud-example/internal/auth/auth_services"
)

type authControllerImpl struct {
	authService authservices.AuthService
}

func SetupAuthController(c *gin.RouterGroup, authService authservices.AuthService) {
	ctrl := authControllerImpl{
		authService: authService,
	}

	c.POST("/login", ctrl.Login)
}

func (a authControllerImpl) Login(c *gin.Context) {
	userCredential := authdto.Login{}
	if err := c.ShouldBindBodyWithJSON(&userCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	fmt.Println("Hello")
	signed, err := a.authService.Verify(c.Request.Context(), 1)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": signed,
	})
}
