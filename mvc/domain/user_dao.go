package domain

import (
	"fmt"
	"github.com/pablogtzgileta/go-microservices-intro/mvc/utils"
	"log"
	"net/http"
)

type userDao struct{}
type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Pablo",
			LastName:  "Gutierrez",
			Email:     "pablo@gmail.com",
		},
	}

	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

func (*userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("were accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
