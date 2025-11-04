package demo

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	authcontrollers "github.com/lunarisnia/crud-example/internal/auth/auth_controllers"
	authservices "github.com/lunarisnia/crud-example/internal/auth/auth_services"
	"github.com/lunarisnia/crud-example/internal/database"
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
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = authservices.FetchToken()
	if err != nil {
		panic(err)
	}

	userRepo := userrepositories.NewUserRepository(db)
	userService := userservices.NewUserService(userRepo)
	authService := authservices.NewAuthService(userService)

	router := server.NewRouter()
	server := &http.Server{
		Addr:    ":3210",
		Handler: router.Handler(),
	}

	baseGroup := router.Group("/v1")
	usercontrollers.SetupUserController(baseGroup, userService)
	authcontrollers.SetupAuthController(baseGroup, authService)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting Down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Println("Database Shutdown:", err)
	}
	log.Println("Server exiting")
}
