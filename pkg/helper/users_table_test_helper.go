package helper

import (
	"day9-cicd/database"
	"day9-cicd/internal/model"
)

func FindUsers() ([]model.User, error) {
	db := database.GetConnection()

	var users []model.User

	err := db.Model(&model.User{}).Find(&users).Error

	return users, err
}

func CleanUsersTable() {
	db := database.GetConnection()

	db.Exec("DELETE FROM users")
}
