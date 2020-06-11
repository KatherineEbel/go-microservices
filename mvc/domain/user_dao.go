package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KatherineEbel/go-microservices/mvc/utils"
)

type UsersServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Katherine",
			LastName:  "Ebel",
			Email:     "kathyebel@me.com",
		},
	}
	UserDao UsersServiceInterface
)

func (u *userDao) GetUser(id int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[id]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", id),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
