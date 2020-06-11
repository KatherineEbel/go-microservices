package services

import (
	"github.com/KatherineEbel/go-microservices/mvc/domain"
	"github.com/KatherineEbel/go-microservices/mvc/utils"
)

type userService struct {
}

var (
	UsersService userService
)

func (u *userService) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(id)
}
