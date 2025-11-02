package demo

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	userRepo := userrepositories.NewUserRepository(db)
	userService := userservices.NewUserService(userRepo)

	s := server.NewRawServer()

	s.Use(gin.Recovery())
	s.Use(middlewares.RequestLogger())
	baseGroup := s.Group("/v1")
	usercontrollers.SetupUserController(baseGroup, userService)

	s.Run("0.0.0.0:3210")
}
