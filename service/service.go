package service

import (
	"context"
	"user_api/model"
	"user_api/repositories"
)

// UserService is a service for users
type UserService interface {
	GetUserById(context.Context, int) (*model.User, error)
	GetAllUsers(context.Context) ([]model.User, error)
	CreateUser(context.Context, string) (*model.User, error)
	UpdateUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(context.Context, int) error
}

type Service struct {
	User UserService
}

func NewService(ctx context.Context, repo *repositories.Repository) *Service {
	return &Service{
		User: NewUserWebService(ctx, repo),
	}
}
