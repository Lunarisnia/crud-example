package userservices

import (
	"context"
	"errors"
	"testing"

	userrepositories_mock "github.com/lunarisnia/crud-example/mocks/users/user_repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	t.Run("Fail", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		userRepoMock := userrepositories_mock.NewMockUserRepository(ctrl)
		defer ctrl.Finish()

		srv := NewUserService(userRepoMock)
		userRepoMock.EXPECT().Create(ctx, nil, gomock.Eq("Test")).Return(errors.New("Err"))

		err := srv.CreateUser(ctx, "Test")
		assert.Error(t, err, "Err")
	})

	t.Run("Success", func(t *testing.T) {
		ctrl, ctx := gomock.WithContext(context.Background(), t)
		userRepoMock := userrepositories_mock.NewMockUserRepository(ctrl)
		defer ctrl.Finish()

		srv := NewUserService(userRepoMock)
		userRepoMock.EXPECT().Create(ctx, nil, gomock.Eq("Test")).Return(nil)

		err := srv.CreateUser(ctx, "Test")
		assert.Equal(t, nil, err)
	})
}
