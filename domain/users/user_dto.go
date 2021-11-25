package users

import (
	"github.com/VoPhan/bookstore_users-api/utils/erros"
	"strings"
)

type User struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	DateCreate string `json:"date_create"`
}

func (user *User) Validate() *erros.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return erros.NewBadRequestError("invalid email address")
	}
	return nil
}
