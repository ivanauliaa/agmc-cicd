package factory

import (
	"day9-cicd/database"
	"day9-cicd/internal/repository"
	"day9-cicd/mongo"
)

type Factory struct {
	UserRepository repository.UserInterface
	BookRepository repository.BookInterface
}

type FactoryMock struct {
	MockUserRepository repository.UserInterface
}

func NewFactory() *Factory {
	db := database.GetConnection()
	mongo := mongo.GetConnection()

	return &Factory{
		UserRepository: repository.NewUser(db),
		BookRepository: repository.NewBook(mongo),
	}
}

func NewFactoryMock(mockUserRepository repository.UserInterface) *Factory {
	return &Factory{
		UserRepository: mockUserRepository,
	}
}
