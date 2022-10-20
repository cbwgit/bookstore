package services

import (
	"bookstore/domain/users"
	"bookstore/utils/crypto_utils"
	"bookstore/utils/date_utils"
	"bookstore/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr){
	dao:= &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}


func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password =crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err!= nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

} 

func DeleteUser(userId int64) *errors.RestErr {
	dao := &users.User{Id: userId}
	return dao.Delete()
}

func Search(status string) ([]users.User , *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
