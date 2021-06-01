package jsonobject

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
	"user_api/config"
	"user_api/model"
)

// TODO: избавиться от байтов и перевести в стракт!!!!!!!!
// UserJson user
type UserJson struct {
	bytes *[]byte
}

// NewUserJson creates new json user
func NewUserJson(bytes *[]byte) *UserJson {
	return &UserJson{bytes: bytes}
}

// CreateUser creates a user, return id
func (r *UserJson) CreateUser(name string) (int, error) {
	log.Printf("json: Create user start")
	defer log.Printf("json: Create user end")

	var (
		id    int
		users model.Users
		err   error
	)

	// unmarshal bytes to struct
	if err := json.Unmarshal(*r.bytes, &users); err != nil {
		log.Printf("CreateUser: unmarshal bytes: %v", err)
		return id, err
	}

	// check user name
	if err := checkUserName(name, users); err != nil {
		return 0, err
	}

	// get new id
	id = users.Users[len(users.Users)-1].ID + 1

	// create a new user
	newUser := model.User{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}

	users.Users = append(users.Users, newUser)

	*r.bytes, err = json.Marshal(users)
	if err != nil {
		log.Printf("CreateUser: marshal bytes: %v", err)
		return 0, err
	}

	// write users to a file
	if err := ioutil.WriteFile(config.Get().JsonPath, *r.bytes, 0660); err != nil {
		log.Printf("CreateUser: write file: %v", err)
		return 0, err
	}

	return id, nil
}

// GetUserById returns user by id from jsonObj
func (r *UserJson) GetUserById(id int) (*model.User, error) {
	log.Printf("json: Get user start. Id: %d", id)
	defer log.Printf("json: Get user end")

	var (
		users   model.Users
		reqUser model.User
	)
	// unmarshal bytes to struct
	if err := json.Unmarshal(*r.bytes, &users); err != nil {
		log.Printf("GetUserById: unmarshal bytes: %v", err)
		return &model.User{}, err
	}

	// find user in users by id
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].ID == id {
			reqUser = users.Users[i]
			break
		}
	}

	return &reqUser, nil
}

// GetAllUsers returns all users from jsonObj in order by id
func (r *UserJson) GetAllUsers() ([]model.User, error) {
	log.Printf("json: Get all users start")
	defer log.Printf("json: Get all users end")

	var users model.Users

	// unmarshal bytes to struct
	if err := json.Unmarshal(*r.bytes, &users); err != nil {
		log.Printf("GetAllUsers: unmarshal bytes: %v", err)
		return []model.User{}, err
	}

	return users.Users, nil
}

// UpdateUser updates user by id
func (r *UserJson) UpdateUser(id int, newName string) error {
	log.Printf("json: Update user start")
	defer log.Printf("json: Update user end")

	var (
		users        model.Users
		reqUserIndex int
	)

	// unmarshal bytes to struct
	if err := json.Unmarshal(*r.bytes, &users); err != nil {
		log.Printf("UpdateUser: unmarshal bytes: %v", err)
		return err
	}

	// check if name is unique
	if err := checkUserName(newName, users); err != nil {
		log.Printf("UpdateUser: incorrect user name: %v", err)
		return err
	}

	// TODO: проверка на то что индекс найден иначе будет менять 0 юзера
	// find a req user index by id
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].ID == id {
			reqUserIndex = i
			break
		}
	}

	users.Users[reqUserIndex].Name = newName

	var err error

	*r.bytes, err = json.Marshal(users)
	if err != nil {
		log.Printf("UpdateUser: marshal bytes: %v", err)
		return err
	}

	// write users to a file
	if err := ioutil.WriteFile(config.Get().JsonPath, *r.bytes, 0660); err != nil {
		log.Printf("UpdateUser: write file: %v", err)
		return err
	}

	return nil
}

// DeleteUser deletes user by id
func (r *UserJson) DeleteUser(id int) error {
	log.Printf("json: Delete user start")
	defer log.Printf("json: Delete user end")

	var (
		users        model.Users
		reqUserIndex int
		err          error
	)

	// unmarshal bytes to struct
	if err := json.Unmarshal(*r.bytes, &users); err != nil {
		log.Printf("DeleteUser: unmarshal bytes: %v", err)
		return err
	}

	// TODO: проверка на то что индекс найден иначе будет менять 0 юзера
	// find req user index by id
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].ID == id {
			reqUserIndex = i
			break
		}
	}

	// delete user
	users.Users = append(users.Users[:reqUserIndex], users.Users[reqUserIndex+1:]...)

	*r.bytes, err = json.Marshal(users)
	if err != nil {
		log.Printf("DeleteUser: marshal bytes: %v", err)
		return err
	}

	// write users to a file
	if err := ioutil.WriteFile(config.Get().JsonPath, *r.bytes, 0660); err != nil {
		log.Printf("DeleteUser: write file: %v", err)
		return err
	}

	return nil
}
