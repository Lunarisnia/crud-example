package users

import (
	"context"
	"database/sql"

	"github.com/lunarisnia/crud-example/internal/database"
)

type UserRepository interface {
	Create(ctx context.Context, tx database.DB, name string) error
}

type userRepositoryImpl struct {
	db database.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u userRepositoryImpl) Create(ctx context.Context, tx database.DB, name string) error {
	if tx != nil {
		u.db = tx
	}

	statement, err := u.db.PrepareContext(ctx,
		"INSERT INTO public.user (name) values ($1)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.ExecContext(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
