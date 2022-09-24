package repository_test

import (
	"context"
	"os"
	"testing"

	"day9-cicd/database"
	"day9-cicd/database/migration"
	"day9-cicd/internal/model"
	"day9-cicd/internal/repository"
	"day9-cicd/pkg/helper"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("APP_ENV", "test")
	godotenv.Load("../../.env")
	migration.Migrate()
}

func TestUser_Create(t *testing.T) {
	db := database.GetConnection()

	t.Run("should not return error", func(t *testing.T) {
		t.Cleanup(func() {
			helper.CleanUsersTable()
		})

		user := model.User{
			Name:     "user",
			Password: "password",
			Email:    "user@gmail.com",
		}
		user.ID = 99

		userRepo := repository.NewUser(db)

		err := userRepo.Create(context.Background(), user)
		assert.NoError(t, err)
	})

	t.Run("should persist data on database", func(t *testing.T) {
		t.Cleanup(func() {
			helper.CleanUsersTable()
		})

		user := model.User{
			Name:     "user",
			Password: "password",
			Email:    "user@gmail.com",
		}
		user.ID = 99

		userRepo := repository.NewUser(db)

		userRepo.Create(context.Background(), user)

		result, _ := helper.FindUsers()
		assert.Equal(t, 1, len(result))
	})
}
