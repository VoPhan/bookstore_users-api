package users

import (
	"fmt"
	"github.com/VoPhan/bookstore_users-api/utils/erros"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *erros.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return erros.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreate = result.DateCreate

	return nil
}

func (user *User) Save() *erros.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return erros.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return erros.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
