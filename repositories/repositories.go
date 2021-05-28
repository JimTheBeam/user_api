package repositories

import (
	"database/sql"
	"user_api/model"
	"user_api/repositories/pg"
)

// UserRepo is a repository for users
type UserRepo interface {
	CreateUser(name string) (int, error)
	GetAllUsers() ([]model.User, error)
	GetUserById(id int) (*model.User, error)
	UpdateUser(id int, name string) error
	DeleteUser(id int) error
}

type Repository struct {
	UserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: pg.NewUserPostgres(db),
	}
}
