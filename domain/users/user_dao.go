package users

import (
	"bookstore/utils/date_utils"
	"bookstore/utils/errors"
	"fmt"
	"bookstore/datasources/mysql/users_db"

)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil { //database connection fail 
		panic(err)
	}
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already register", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))

	}
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
