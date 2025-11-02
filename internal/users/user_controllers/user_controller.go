package usercontrollers

import (
	"net/http"
	"strconv"

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
	r.GET("/user/:id", c.GetUserById)
	r.POST("/user", c.CreateUser)
	r.PATCH("/user/:id", c.UpdateName)
	r.DELETE("/user/:id", c.DeleteUser)
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

func (u *userControllerImpl) GetUserById(c *gin.Context) {
	userIdStr := c.Param("id")
	if userIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	user, err := u.userService.GetUserById(c.Request.Context(), uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *userControllerImpl) UpdateName(c *gin.Context) {
	userIdStr := c.Param("id")
	if userIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	body := userdto.UpdateName{}
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err = u.userService.UpdateName(c.Request.Context(), uint(userIdInt), body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Name Updated",
	})
}

func (u *userControllerImpl) DeleteUser(c *gin.Context) {
	userIdStr := c.Param("id")
	if userIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	err = u.userService.DeleteUser(c.Request.Context(), uint(userIdInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Removed",
	})
}
