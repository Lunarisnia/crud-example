package demo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lunarisnia/crud-example/internal/database"
	"github.com/lunarisnia/crud-example/internal/users"
)

func Run() {
	ctx := context.Background()
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatalln("DATABASE_URL is empty")
	}
	db, err := database.Connect(databaseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	userRepo := users.NewUserRepository(db)
	createUserDemo(ctx, db, userRepo)

	// With wrapper
	database.UseTransaction(ctx, db, func(tx *sql.Tx) error {
		err = userRepo.Create(ctx, tx, "This")
		if err != nil {
			log.Fatalln(err)
		}
		err = userRepo.Create(ctx, tx, "Automatically, Rollback and Commit")
		if err != nil {
			log.Fatalln()
		}
		return nil
	})

	user, err := userRepo.ReadById(ctx, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
}

func createUserDemo(ctx context.Context, db *sql.DB, userRepo users.UserRepository) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = userRepo.Create(ctx, tx, "Foobar")
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}

	err = userRepo.Create(ctx, tx, "Hello, World")
	if err != nil {
		tx.Rollback()
		log.Fatalln()
	}
	tx.Commit()
}
