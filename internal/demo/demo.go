package demo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lunarisnia/crud-example/internal/database"
	"github.com/lunarisnia/crud-example/internal/users"
	"gorm.io/gorm"
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

	userRepo := users.NewUserRepository(db)

	// NOTE: Create User
	createUserDemo(ctx, db, userRepo)
	// NOTE: Create User With Transaction wrapper
	database.UseTransaction(ctx, db, func(tx *gorm.DB) error {
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

	// NOTE: Read user by ID
	// showFirstUser(ctx, userRepo)

	// NOTE: Read all User
	// showAllUser(ctx, userRepo)

	// NOTE: Change first user name
	// changeFirstUserName(ctx, userRepo)
	// showFirstUser(ctx, userRepo)

	// NOTE: Delete the last user
	// removeLastUser(ctx, userRepo)
}

func createUserDemo(ctx context.Context, db *gorm.DB, userRepo users.UserRepository) {
	tx := db.Begin()

	err := userRepo.Create(ctx, tx, "Foobar")
	if err != nil {
		tx.Rollback()
		log.Fatalln(err)
	}
	tx.Commit()
}

func showFirstUser(ctx context.Context, userRepo users.UserRepository) {
	user, err := userRepo.ReadById(ctx, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("===========ReadByID==========")
	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("===========ReadByID==========")
}

func showAllUser(ctx context.Context, userRepo users.UserRepository) {
	users, err := userRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, u := range users {
		fmt.Println("=====================")
		fmt.Println("ID:", u.ID)
		fmt.Println("Name:", u.Name)
	}
}

func changeFirstUserName(ctx context.Context, userRepo users.UserRepository) {
	err := userRepo.UpdateName(ctx, nil, 1, "Training")
	if err != nil {
		log.Fatalln(err)
	}
}

func removeLastUser(ctx context.Context, userRepo users.UserRepository) {
	users, err := userRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Total User (Before Deletion):", len(users))

	err = userRepo.Remove(ctx, nil, users[len(users)-1].ID)
	if err != nil {
		log.Fatalln(err)
	}

	users, err = userRepo.ReadAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Total User (After Deletion):", len(users))
}
