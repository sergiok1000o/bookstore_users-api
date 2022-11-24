package services

//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
import (
	"github.com/sergio/bookstore_users-api/domain/users"
	"github.com/sergio/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}