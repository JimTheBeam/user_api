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
func (svc *UserWebService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := svc.repo.GetAllUsers()
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetAllUsers")
	}

	return users, nil
}

// CreateUser ...
func (svc *UserWebService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	// create a new user
	createdID, err := svc.repo.CreateUser(name)
	if err != nil {
		return nil, errors.Wrap(err, "svc.repo.CreateUser error")
	}

	// get created user by ID
	createdDBUser, err := svc.repo.GetUserById(createdID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.repo.GetUser error")
	}

	return createdDBUser, nil
}

// UpdateUser ...
func (svc *UserWebService) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// update user
	err := svc.repo.UpdateUser(user.ID, user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.UpdateUser")
	}

	// get updated user by ID
	updatedUser, err := svc.repo.GetUserById(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "svc.repo.GetUser error")
	}

	return updatedUser, nil
}

// DeleteUser ...
func (svc *UserWebService) DeleteUser(ctx context.Context, userID int) error {
	err := svc.repo.DeleteUser(userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.DeleteUser")
	}

	return nil
}
