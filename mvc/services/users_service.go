package services

import (
	"github.com/pablogtzgileta/go-microservices-intro/mvc/domain"
	"github.com/pablogtzgileta/go-microservices-intro/mvc/utils"
)

type usersService struct{}

var UsersService usersService

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
