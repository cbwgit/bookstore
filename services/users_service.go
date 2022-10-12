package services

import (
	"bookstore/domain/users"
	"bookstore/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err!= nil{
		return nil, err
	}
	return &user, nil
}
