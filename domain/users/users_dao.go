package users

import (
	"fmt"

	"github.com/sergio/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LasttName = result.LasttName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	
	usersDB[user.Id] = &user
	return nil
}

