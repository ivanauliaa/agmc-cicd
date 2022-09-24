package user_test

import (
	"context"
	"testing"

	"day9-cicd/internal/app/user"
	"day9-cicd/internal/dto"
	"day9-cicd/internal/factory"
	"day9-cicd/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Create(t *testing.T) {
	t.Run("should orchestrating create user service", func(t *testing.T) {
		mockUserRepository := repository.NewMockUser()
		mockUserRepository.On("Create", mock.Anything, mock.Anything).Return(nil)

		factory := factory.NewFactoryMock(mockUserRepository)

		svc := user.NewService(factory)

		payload := dto.CreateUserRequest{
			Name:     "user",
			Email:    "user@gmail.com",
			Password: "password",
		}

		result, err := svc.Create(context.Background(), &payload)

		assert.NoError(t, err)
		assert.Equal(t, "success", result)
	})
}
