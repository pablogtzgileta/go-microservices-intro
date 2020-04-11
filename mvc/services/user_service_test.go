package services

import (
	"github.com/pablogtzgileta/go-microservices-intro/mvc/domain"
	"github.com/pablogtzgileta/go-microservices-intro/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &userDaoMock
}

type usersDaoMock struct{}

func (n *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestUsersNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
		}
	}

	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
