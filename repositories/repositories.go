package repositories

import (
	"database/sql"
	"user_api/model"
	jsonobject "user_api/repositories/jsonObject"
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

// NewRepository creates new database repository
func NewRepositoryDB(db *sql.DB) *Repository {
	return &Repository{
		UserRepo: pg.NewUserPostgres(db),
	}
}

// NewRepository creates new json repository
func NewRepositoryJson(js *model.Users) *Repository {
	return &Repository{
		UserRepo: jsonobject.NewUserJson(js),
	}
}
