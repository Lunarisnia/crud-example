package userservices

import (
	"context"

	userentities "github.com/lunarisnia/crud-example/internal/users/user_entities"
	userrepositories "github.com/lunarisnia/crud-example/internal/users/user_repositories"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]userentities.User, error)
	CreateUser(ctx context.Context, name string) error
}

type userServiceImpl struct {
	userRepo userrepositories.UserRepository
}

func NewUserService(userRepo userrepositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (u *userServiceImpl) GetAllUser(ctx context.Context) ([]userentities.User, error) {
	users, err := u.userRepo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userServiceImpl) CreateUser(ctx context.Context, name string) error {
	err := u.userRepo.Create(ctx, nil, name)
	if err != nil {
		return err
	}

	return nil
}
