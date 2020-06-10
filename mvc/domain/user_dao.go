package domain

import (
	"fmt"
	"github.com/KatherineEbel/go-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Katherine",
			LastName:  "Ebel",
			Email:     "kathyebel@me.com",
		},
	}
)

func GetUser(id int64) (*User, *utils.ApplicationError) {
	if user := users[id]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", id),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
