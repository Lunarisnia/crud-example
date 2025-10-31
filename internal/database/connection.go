package database

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func UseTransaction(ctx context.Context, db *sql.DB, action func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = action(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
