package usercontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userdto "github.com/lunarisnia/crud-example/internal/users/user_dto"
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
	r.POST("/user", c.CreateUser)
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

func (u *userControllerImpl) CreateUser(c *gin.Context) {
	body := userdto.CreateUser{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := u.userService.CreateUser(c.Request.Context(), body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
