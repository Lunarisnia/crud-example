package database

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(databaseUrl string) (*gorm.DB, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func UseTransaction(ctx context.Context, db *gorm.DB, action func(tx *gorm.DB) error) error {
	tx := db.Begin()

	err := action(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
