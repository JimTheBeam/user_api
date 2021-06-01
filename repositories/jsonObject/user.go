package jsonobject

import (
	"errors"
	"fmt"
	"log"
	"time"
	"user_api/model"
)

// UserJson user
type UserJson struct {
	userJS *model.Users
}

// NewUserJson creates new json user
func NewUserJson(users *model.Users) *UserJson {
	return &UserJson{userJS: users}
}

// CreateUser creates a user, return id
func (r *UserJson) CreateUser(name string) (int, error) {
	log.Printf("json: Create user start")
	defer log.Printf("json: Create user end")

	id := 1

	// check user name
	if err := checkUserName(name, r.userJS); err != nil {
		return 0, err
	}

	// get new id
	if len(r.userJS.Users) > 0 {
		id = r.userJS.Users[len(r.userJS.Users)-1].ID + 1
	}

	// create a new user
	newUser := model.User{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05"),
	}

	r.userJS.Users = append(r.userJS.Users, newUser)

	if err := writeJsonFile(r.userJS); err != nil {
		log.Printf("CreateUser write file: %v", err)
		return 0, err
	}

	return id, nil
}

// GetUserById returns user by id from jsonObj
func (r *UserJson) GetUserById(id int) (*model.User, error) {
	log.Printf("json: Get user start. Id: %d", id)
	defer log.Printf("json: Get user end")

	var (
		reqUser model.User
		exist   bool
	)

	// find user in users by id
	for i := 0; i < len(r.userJS.Users); i++ {
		if r.userJS.Users[i].ID == id {
			reqUser = r.userJS.Users[i]
			exist = true
			break
		}
	}

	// check if id exists
	if !exist {
		errString := fmt.Sprintf("user with id=%v not found", id)
		log.Printf(errString)
		return &model.User{}, errors.New(errString)
	}

	return &reqUser, nil
}

// GetAllUsers returns all users from jsonObj in order by id
func (r *UserJson) GetAllUsers() ([]model.User, error) {
	log.Printf("json: Get all users start")
	defer log.Printf("json: Get all users end")

	return r.userJS.Users, nil
}

// UpdateUser updates user by id
func (r *UserJson) UpdateUser(id int, newName string) error {
	log.Printf("json: Update user start")
	defer log.Printf("json: Update user end")

	var (
		reqUserIndex int
		exist        bool
	)

	// check if name is unique
	if err := checkUserName(newName, r.userJS); err != nil {
		log.Printf("UpdateUser: incorrect user name: %v", err)
		return err
	}

	// find a req user index by id
	for i := 0; i < len(r.userJS.Users); i++ {
		if r.userJS.Users[i].ID == id {
			reqUserIndex = i
			exist = true
			break
		}
	}

	// check if user exists
	if !exist {
		errString := fmt.Sprintf("user with id=%v not found", id)
		log.Printf(errString)
		return errors.New(errString)
	}

	r.userJS.Users[reqUserIndex].Name = newName

	if err := writeJsonFile(r.userJS); err != nil {
		log.Printf("UpdateUser write file: %v", err)
		return err
	}

	return nil
}

// DeleteUser deletes user by id
func (r *UserJson) DeleteUser(id int) error {
	log.Printf("json: Delete user start")
	defer log.Printf("json: Delete user end")

	var (
		reqUserIndex int
		exist        bool
	)

	// find req user index by id
	for i := 0; i < len(r.userJS.Users); i++ {
		if r.userJS.Users[i].ID == id {
			reqUserIndex = i
			exist = true
			break
		}
	}

	// check if user exists
	if !exist {
		errString := fmt.Sprintf("user with id=%v not found", id)
		log.Printf(errString)
		return errors.New(errString)
	}

	// delete user
	r.userJS.Users = append(r.userJS.Users[:reqUserIndex], r.userJS.Users[reqUserIndex+1:]...)

	if err := writeJsonFile(r.userJS); err != nil {
		log.Printf("DeleteUser write file: %v", err)
		return err
	}

	return nil
}
