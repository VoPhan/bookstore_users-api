package services

import (
	"github.com/VoPhan/bookstore_users-api/domain/users"
	"github.com/VoPhan/bookstore_users-api/utils/erros"
)

func GerUser(userId int64) (*users.User, *erros.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *erros.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
