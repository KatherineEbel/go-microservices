package services

import (
	"net/http"

	"github.com/KatherineEbel/go-microservices/mvc/domain"
	"github.com/KatherineEbel/go-microservices/mvc/utils"
)

type itemService struct {
}

var (
	ItemsService itemService
)

func (i itemService) GetItem(id string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "internal_server_error",
	}
}
