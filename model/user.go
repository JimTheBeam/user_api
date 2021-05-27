package model

import (
	"time"
)

// User is a json user
type User struct {
	ID        int       `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
