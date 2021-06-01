package jsonobject

import (
	"errors"
	"fmt"
	"user_api/model"
)

// checkUserName checks if username unique
func checkUserName(name string, users *model.Users) error {
	for i := 0; i < len(users.Users); i++ {
		if name == users.Users[i].Name {
			return errors.New(fmt.Sprintf("user %s already exists.", name))
		}
	}
	return nil
}
