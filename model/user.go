package model

import (
	"time"
)

// User is a json user
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// Users is a json users
type Users struct {
	Users []User `json:"users" validate:"required"`
}
