package demo

import (
	"log"
	"os"

	"github.com/lunarisnia/crud-example/internal/database"
	"github.com/lunarisnia/crud-example/internal/server"
	"github.com/lunarisnia/crud-example/internal/users"
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

	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)

	s := server.NewServer()
	baseGroup := s.Group("/v1")
	users.SetupUserController(baseGroup, userService)
	s.Run("0.0.0.0:3210")
}
