package usercontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userservices "github.com/lunarisnia/crud-example/internal/users/user_services"
)

type userControllerImpl struct {
	userService userservices.UserService
}

func SetupUserController(r *gin.RouterGroup, userService userservices.UserService) {
	c := userControllerImpl{
		userService: userService,
	}
	r.GET("/user", c.GetAllUser)
}

func (u *userControllerImpl) GetAllUser(c *gin.Context) {
	users, err := u.userService.GetAllUser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
