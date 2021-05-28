package service

import (
	"context"
	"fmt"
	"user_api/lib/types"
	"user_api/model"
	"user_api/repositories"

	"github.com/pkg/errors"
)

// UserWebService ...
type UserWebService struct {
	repo *repositories.Repository
	ctx  context.Context
}

// NewUserWebService creates a new user web service
func NewUserWebService(ctx context.Context, repo *repositories.Repository) *UserWebService {
	return &UserWebService{
		repo: repo,
		ctx:  ctx,
	}
}

// GetUserById ...
func (svc *UserWebService) GetUserById(ctx context.Context, userID int) (*model.User, error) {
	userDB, err := svc.repo.GetUserById(userID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetUser")
	}
	if userDB == nil {
		return nil, errors.Wrap(types.ErrNotFound, fmt.Sprintf("User '%v' not found", userID))
	}

	return userDB, nil
}

// GetAllUsers ...
func (svc *UserWebService) GetAllUsers(ctx context.Context, user *model.User) ([]model.User, error) {
	return nil, nil
}

// CreateUser ...
func (svc *UserWebService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

// UpdateUser ...
func (svc *UserWebService) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

// DeleteUser ...
func (svc *UserWebService) DeleteUser(ctx context.Context, userID int) error {
	return nil
}
