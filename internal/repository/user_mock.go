package repository

import (
	"context"
	"day9-cicd/internal/model"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type mockUser struct {
	mock.Mock
}

func NewMockUser() *mockUser {
	return &mockUser{
		mock.Mock{},
	}
}

func (r *mockUser) Create(ctx context.Context, data model.User) error {
	fmt.Println("mock user create", data)
	args := r.Called(ctx, data)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(error)
}

func (r *mockUser) FindAll(ctx context.Context) ([]model.User, error) {
	args := r.Called(ctx)
	return args.Get(0).([]model.User), args.Get(1).(error)
}

func (r *mockUser) FindById(ctx context.Context, id uint) (model.User, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(model.User), args.Get(1).(error)
}

func (r *mockUser) Update(ctx context.Context, id uint, data map[string]interface{}) error {
	args := r.Called(ctx, id, data)
	return args.Get(0).(error)
}

func (r *mockUser) Delete(ctx context.Context, id uint) error {
	args := r.Called(ctx, id)
	return args.Get(0).(error)
}

func (r *mockUser) FindByEmail(ctx context.Context, email string) (model.User, error) {
	args := r.Called(ctx, email)
	return args.Get(0).(model.User), args.Get(1).(error)
}

func (r *mockUser) VerifyUserOwner(ctx context.Context, id uint, owner uint) error {
	args := r.Called(ctx, id, owner)
	return args.Get(0).(error)
}
