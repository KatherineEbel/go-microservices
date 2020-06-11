package services

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KatherineEbel/go-microservices/mvc/domain"
	"github.com/KatherineEbel/go-microservices/mvc/utils"
)

type usersDaoMock struct {
}

var (
	getUserFunction func(id int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(id)
}

func TestUserService_GetUserNotFound(t *testing.T) {
	getUserFunction = func(id int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exist",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.NotNil(t, err.Message)
}

func TestUserService_GetUserFound(t *testing.T) {
	getUserFunction = func(id int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        0,
			FirstName: "john",
			LastName:  "doe",
			Email:     "johndoe@example.com",
		}, nil
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 0, user.Id)
}
