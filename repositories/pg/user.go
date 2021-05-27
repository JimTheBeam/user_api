package pg

import (
	"database/sql"
	"fmt"
	"log"
	"user_api/model"
)

// UserPgPostgres user
type UserPgPostgres struct {
	db *sql.DB
}

// NewUserPostgres create new
func NewUserPostgres(db *sql.DB) *UserPgPostgres {
	return &UserPgPostgres{db: db}
}

// CreateUser creates user in Postgres return id
func (r *UserPgPostgres) CreateUser(name string) (int, error) {
	log.Printf("DB: Create user start")
	defer log.Printf("DB: Create user end")

	var id int

	sql := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", usersTable)

	if err := r.db.QueryRow(sql, name).Scan(&id); err != nil {
		log.Printf("DB: Create query: %v", err)
		return 0, err
	}

	return id, nil
}

// GetUserById returns user by id from db
func (r *UserPgPostgres) GetUserById(id int) (model.User, error) {
	log.Printf("DB: Get user start")
	defer log.Printf("DB: Get user end")

	var user model.User

	sql := fmt.Sprintf("SELECT id, name, created_at FROM %s WHERE id = $1", usersTable)

	err := r.db.QueryRow(sql, id).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		log.Printf("DB: Get query: %v", err)
		return model.User{}, err
	}

	return user, nil
}

// GetAllUsers returns all users from db in order by id
func (r *UserPgPostgres) GetAllUsers() ([]model.User, error) {
	log.Printf("DB: Get all users start")
	defer log.Printf("DB: Get all users end")

	var users []model.User

	sql := fmt.Sprintf("SELECT id, name, created_at FROM %s ORDER BY id", usersTable)

	rows, err := r.db.Query(sql)
	if err != nil {
		log.Printf("DB: Get all query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser updates user by id
func (r *UserPgPostgres) UpdateUser(id int, name string) error {
	log.Printf("DB: Update user start")
	defer log.Printf("DB: Update user end")

	sql := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", usersTable)

	_, err := r.db.Exec(sql, name, id)
	if err != nil {
		log.Printf("DB: Update query: %v", err)
		return err
	}

	return nil
}

// DeleteUser deletes user by id
func (r *UserPgPostgres) DeleteUser(id int) error {
	log.Printf("DB: Delete user start")
	defer log.Printf("DB: Delete user end")

	sql := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)

	_, err := r.db.Exec(sql, id)
	if err != nil {
		log.Printf("DB: Delete query: %v", err)
		return err
	}
	return nil
}
