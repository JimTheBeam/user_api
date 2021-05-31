package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"user_api/lib/types"
	"user_api/model"
	"user_api/service"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// UserHandler
type UserHandler struct {
	service *service.Service
}

// NewUsers creates a new user handler
func NewUsers(ctx context.Context, service *service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Create creates new user
func (hnd *UserHandler) Create(ctx echo.Context) error {
	log.Printf("Create handler starts")

	var jsonUser model.User

	if err := ctx.Bind(&jsonUser); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not create user"))
	}

	if err := ctx.Validate(&jsonUser); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	log.Println("Create user struct:", jsonUser)

	createdUser, err := hnd.service.User.CreateUser(jsonUser.Name)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create user"))
		}
	}

	return ctx.JSON(http.StatusCreated, createdUser)
}

// Get returns user by ID
func (hnd *UserHandler) GetUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("could not parse user ID")
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user ID"))
	}
	log.Printf("Get user handler ID: %v", userID)

	// get user
	user, err := hnd.service.User.GetUserById(userID)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get user"))
		}
	}
	return ctx.JSON(http.StatusOK, user)
}

// GetAllUsers returns user by ID
func (hnd *UserHandler) GetAllUsers(ctx echo.Context) error {
	log.Printf("Get all users handler")

	users, err := hnd.service.User.GetAllUsers()
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get all users"))
	}
	return ctx.JSON(http.StatusOK, users)
}

// DeleteUser returns user by ID
func (hnd *UserHandler) DeleteUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("could not parse user ID")
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user ID"))
	}
	log.Printf("Delete user handler ID: %v", userID)

	// delete user
	err = hnd.service.User.DeleteUser(userID)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get user"))
		}
	}
	return echo.NewHTTPError(http.StatusOK)

}

// Update updates user in DB
func (hnd *UserHandler) Update(ctx echo.Context) error {
	log.Printf("Update handler starts")

	var reqUser model.User

	if err := ctx.Bind(&reqUser); err != nil {
		log.Printf("Bind json Error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not create user"))
	}

	if err := ctx.Validate(&reqUser); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	log.Printf("Update handler. User: %v", reqUser)

	updatedUser, err := hnd.service.User.UpdateUser(&reqUser)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not update user"))
		}
	}

	return ctx.JSON(http.StatusCreated, updatedUser)
}
