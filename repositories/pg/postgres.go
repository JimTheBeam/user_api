package pg

import (
	"database/sql"
	"fmt"
	"log"
	"user_api/config"

	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

// NewPostgresDB creates connection to postgres database
func NewPostgresDB() (*sql.DB, error) {
	// get config
	cfg := config.Get()

	// connect to database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.DBName,
		cfg.DB.Password,
		cfg.DB.SSLMode,
	),
	)
	if err != nil {
		log.Printf("Database connection: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Lost database connection: %v", err)
		return nil, err
	}
	log.Println("Successfully connected to database.")

	return db, nil
}
