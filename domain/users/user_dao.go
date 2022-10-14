package users

import (
	"bookstore/datasources/mysql/users_db"
	"bookstore/utils/date_utils"
	"bookstore/utils/errors"
	"fmt"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	//result, err:= users_db.Client.Exec(queryInsertUser, user.LastName, user.Email, user.DateCreated)

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewInternalServerError(
				fmt.Sprintf("Email %s already exists!", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: ", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to get last insert ID: %s", err.Error()))
	}
	user.Id = userId
	return nil
}
