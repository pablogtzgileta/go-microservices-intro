package app

import (
	"github.com/pablogtzgileta/go-microservices-intro/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}