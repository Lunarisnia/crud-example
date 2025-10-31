package users

import (
	"context"
	"database/sql"

	"github.com/lunarisnia/crud-example/internal/database"
)

type UserRepository interface {
	Create(ctx context.Context, tx database.DB, name string) error
	ReadById(ctx context.Context, id int) (*User, error)
	ReadAll(ctx context.Context) ([]*User, error)
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

func (u userRepositoryImpl) ReadById(ctx context.Context, id int) (*User, error) {
	statement, err := u.db.PrepareContext(ctx,
		"SELECT * from public.user where id = $1 limit 1")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := User{}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u userRepositoryImpl) ReadAll(ctx context.Context) ([]*User, error) {
	statement, err := u.db.PrepareContext(ctx,
		"SELECT * from public.user")
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
