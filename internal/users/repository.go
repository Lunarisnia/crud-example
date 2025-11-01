package users

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, name string) error
	ReadById(ctx context.Context, id int) (User, error)
	ReadAll(ctx context.Context) ([]User, error)
	UpdateName(ctx context.Context, tx *gorm.DB, id uint, newName string) error
	Remove(ctx context.Context, tx *gorm.DB, id uint) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u userRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, name string) error {
	if tx != nil {
		u.db = tx
	}

	err := gorm.G[User](u.db).Create(ctx, &User{
		Name: name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpl) ReadById(ctx context.Context, id int) (User, error) {
	user, err := gorm.G[User](u.db).First(ctx)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u userRepositoryImpl) ReadAll(ctx context.Context) ([]User, error) {
	users, err := gorm.G[User](u.db).Find(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepositoryImpl) UpdateName(ctx context.Context, tx *gorm.DB, id uint, newName string) error {
	if tx != nil {
		u.db = tx
	}

	_, err := gorm.G[User](u.db).Where("id = ?", id).Update(ctx, "name", newName)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpl) Remove(ctx context.Context, tx *gorm.DB, id uint) error {
	if tx != nil {
		u.db = tx
	}

	_, err := gorm.G[User](u.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
