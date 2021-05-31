package service

import (
	"fmt"
	"user_api/lib/types"
	"user_api/model"
	"user_api/repositories"

	"github.com/pkg/errors"
)

// UserWebService ...
type UserWebService struct {
	repo *repositories.Repository
}

// NewUserWebService creates a new user web service
func NewUserWebService(repo *repositories.Repository) *UserWebService {
	return &UserWebService{
		repo: repo,
	}
}

// GetUserById ...
func (svc *UserWebService) GetUserById(userID int) (*model.User, error) {
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
func (svc *UserWebService) GetAllUsers() ([]model.User, error) {
	users, err := svc.repo.GetAllUsers()
	if err != nil {
		return nil, errors.Wrap(err, "svc.user.GetAllUsers")
	}

	return users, nil
}

// CreateUser ...
func (svc *UserWebService) CreateUser(name string) (*model.User, error) {
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
func (svc *UserWebService) UpdateUser(user *model.User) (*model.User, error) {
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
func (svc *UserWebService) DeleteUser(userID int) error {
	err := svc.repo.DeleteUser(userID)
	if err != nil {
		return errors.Wrap(err, "svc.user.DeleteUser")
	}

	return nil
}
