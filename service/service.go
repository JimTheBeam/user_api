package service

import (
	"user_api/model"
	"user_api/repositories"
)

// UserService is a service for users
type UserService interface {
	GetUserById(int) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	CreateUser(string) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(int) error
}

type Service struct {
	User UserService
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		User: NewUserWebService(repo),
	}
}
