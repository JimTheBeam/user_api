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
	ctx context.Context
	// TODO:
	service *service.Service
}

// NewUsers creates a new user handler
func NewUsers(ctx context.Context, service *service.Service) *UserHandler {
	return &UserHandler{
		ctx:     ctx,
		service: service,
	}
}

// Create TODO: change func
func (hnd *UserHandler) Create(ctx echo.Context) error {
	var user model.User

	err := ctx.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode user data"))
	}

	err = ctx.Validate(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}

	// createdUser, err := hnd.service.User.CreateUser(ctx.Request().Context(), &user)
	// if err != nil {
	// 	switch {
	// 	case errors.Cause(err) == types.ErrBadRequest:
	// 		return echo.NewHTTPError(http.StatusBadRequest, err)
	// 	default:
	// 		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create user"))
	// 	}
	// }

	// return ctx.JSON(http.StatusCreated, createdUser)
	return nil
}

// Get returns user by ID
func (hnd *UserHandler) GetUser(ctx echo.Context) error {

	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user ID"))
	}
	log.Printf("Get user handler ID: %v", userID)

	user, err := hnd.service.User.GetUserById(ctx.Request().Context(), userID)
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
