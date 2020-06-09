package services

import (
	"github.com/KatherineEbel/go-microservices/mvc/domain"
	"github.com/KatherineEbel/go-microservices/mvc/utils"
)

func GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(id)
}
