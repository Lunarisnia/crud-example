package usercontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunarisnia/crud-example/internal/auth"
	"github.com/lunarisnia/crud-example/internal/middlewares"
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
	r.GET("/users", middlewares.VerifyAuth(), c.GetAllUser)
	r.POST("/user", middlewares.VerifyAuth(), c.CreateUser)

	r.GET("/user", middlewares.VerifyAuth(), c.GetUserById)
	r.PATCH("/user", middlewares.VerifyAuth(), c.UpdateName)
	r.DELETE("/user", middlewares.VerifyAuth(), c.DeleteUser)
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
	userCredential, err := auth.ParseJWTClaim(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}

	user, err := u.userService.GetUserById(c.Request.Context(), userCredential.ID)
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
	userCredential, err := auth.ParseJWTClaim(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
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

	err = u.userService.UpdateName(c.Request.Context(), userCredential.ID, body.Name)
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
	userCredential, err := auth.ParseJWTClaim(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}

	err = u.userService.DeleteUser(c.Request.Context(), userCredential.ID)
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
