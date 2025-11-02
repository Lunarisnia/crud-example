package users

import "context"

type UserService interface {
	GetAllUser(ctx context.Context) ([]User, error)
}

type userServiceImpl struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (u *userServiceImpl) GetAllUser(ctx context.Context) ([]User, error) {
	users, err := u.userRepo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
