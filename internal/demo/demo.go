package demo

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	authcontrollers "github.com/lunarisnia/crud-example/internal/auth/auth_controllers"
	authservices "github.com/lunarisnia/crud-example/internal/auth/auth_services"
	"github.com/lunarisnia/crud-example/internal/database"
	"github.com/lunarisnia/crud-example/internal/middlewares"
	"github.com/lunarisnia/crud-example/internal/server"
	usercontrollers "github.com/lunarisnia/crud-example/internal/users/user_controllers"
	userrepositories "github.com/lunarisnia/crud-example/internal/users/user_repositories"
	userservices "github.com/lunarisnia/crud-example/internal/users/user_services"
)

func Run() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatalln("DATABASE_URL is empty")
	}
	db, err := database.Connect(databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	authservices.FetchToken()

	userRepo := userrepositories.NewUserRepository(db)
	userService := userservices.NewUserService(userRepo)
	authService := authservices.NewAuthService(userService)

	s := server.NewRawServer()

	s.Use(gin.Recovery())
	s.Use(middlewares.RequestLogger())
	baseGroup := s.Group("/v1")
	usercontrollers.SetupUserController(baseGroup, userService)
	authcontrollers.SetupAuthController(baseGroup, authService)

	s.Run("0.0.0.0:3210")
}
