package userrepositories

import (
	"context"
	"fmt"
	"log"
	"time"

	userentities "github.com/lunarisnia/crud-example/internal/users/user_entities"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, name string) error
	ReadById(ctx context.Context, id uint) (userentities.User, error)
	ReadAll(ctx context.Context) ([]userentities.User, error)
	UpdateName(ctx context.Context, tx *gorm.DB, id uint, newName string) error
	Remove(ctx context.Context, tx *gorm.DB, id uint) error
}

type userRepositoryImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserRepository(db *gorm.DB, rdb *redis.Client) UserRepository {
	return &userRepositoryImpl{
		db:  db,
		rdb: rdb,
	}
}

func (u userRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, name string) error {
	if tx != nil {
		u.db = tx
	}

	user := userentities.User{
		Name: name,
	}
	err := gorm.G[userentities.User](u.db).Create(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpl) ReadById(ctx context.Context, id uint) (userentities.User, error) {
	var user userentities.User
	err := u.rdb.HGetAll(ctx, fmt.Sprint("user:", id)).Scan(&user)
	if err != nil {
		return user, err
	}
	// Cache Hit! Return the cached result
	if user.ID != 0 {
		log.Println("Cache HIT!:", user.ID)
		return user, nil
	}

	user, err = gorm.G[userentities.User](u.db).First(ctx)
	if err != nil {
		return user, err
	}

	// Write to cache on cache miss
	log.Println("Populating Cache for user:", user.ID)
	_, err = u.rdb.HSet(ctx, fmt.Sprint("user:", user.ID), user).Result()
	if err != nil {
		return user, err
	}
	_, err = u.rdb.Expire(ctx, fmt.Sprint("user:", user.ID), 30*time.Second).Result()
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u userRepositoryImpl) ReadAll(ctx context.Context) ([]userentities.User, error) {
	users, err := gorm.G[userentities.User](u.db).Order("id asc").Find(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepositoryImpl) UpdateName(ctx context.Context, tx *gorm.DB, id uint, newName string) error {
	if tx != nil {
		u.db = tx
	}

	_, err := gorm.G[userentities.User](u.db).Where("id = ?", id).Update(ctx, "name", newName)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpl) Remove(ctx context.Context, tx *gorm.DB, id uint) error {
	if tx != nil {
		u.db = tx
	}

	_, err := gorm.G[userentities.User](u.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
