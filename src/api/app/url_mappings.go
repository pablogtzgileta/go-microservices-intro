package app

import (
	"github.com/pablogtzgileta/go-microservices-intro/src/api/controllers/polo"
	"github.com/pablogtzgileta/go-microservices-intro/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("marco", polo.Polo) // To let know the cloud (Amazon) that we are ready to receive traffic if working
	router.POST("/repositories", repositories.CreateRepo)
}
