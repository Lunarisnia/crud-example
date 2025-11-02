package userservices

import (
	"context"

	userentities "github.com/lunarisnia/crud-example/internal/users/user_entities"
	userrepositories "github.com/lunarisnia/crud-example/internal/users/user_repositories"
)

type UserService interface {
	GetAllUser(ctx context.Context) ([]userentities.User, error)
	GetUserById(ctx context.Context, id uint) (userentities.User, error)
	CreateUser(ctx context.Context, name string) error
	UpdateName(ctx context.Context, id uint, newName string) error
	DeleteUser(ctx context.Context, id uint) error
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
	return u.userRepo.ReadAll(ctx)
}

func (u *userServiceImpl) CreateUser(ctx context.Context, name string) error {
	return u.userRepo.Create(ctx, nil, name)
}

func (u *userServiceImpl) GetUserById(ctx context.Context, id uint) (userentities.User, error) {
	return u.userRepo.ReadById(ctx, id)
}

func (u *userServiceImpl) UpdateName(ctx context.Context, id uint, newName string) error {
	return u.userRepo.UpdateName(ctx, nil, id, newName)
}

func (u *userServiceImpl) DeleteUser(ctx context.Context, id uint) error {
	return u.userRepo.Remove(ctx, nil, id)
}
