package users

import (
	"bookstore/datasources/mysql/users_db"
	"bookstore/utils/date_utils"
	"bookstore/utils/errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no row result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		sqlErr, ok := getErr.(*mysql.MySQLError)
		if !ok {
				
			return errors.NewInternalServerError(fmt.Sprintf("error parsing database response %s",getErr.Error()))
		}
		fmt.Println(sqlErr.Number)
		fmt.Println(sqlErr.Message)

		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprint("user not exist", user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying get user %d:%s", user.Id, err.Error()))
	}
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

	insertResult, saveErr  := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil{
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
				
			return errors.NewInternalServerError(fmt.Sprintf("error parsing database response %s",saveErr.Error()))
		}
	
		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("invalid data",user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error processing request",saveErr.Error()))
	}
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
