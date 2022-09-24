package auth

import (
	"context"
	"day9-cicd/internal/infrastructure/jwt"

	"day9-cicd/internal/dto"
	"day9-cicd/internal/factory"
	"day9-cicd/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	Login(ctx context.Context, payload *dto.UserLoginRequest) (string, error)
}

type service struct {
	UserRepository repository.UserInterface
}

func NewService(f *factory.Factory) ServiceInterface {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Login(ctx context.Context, payload *dto.UserLoginRequest) (string, error) {
	user, err := s.UserRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
