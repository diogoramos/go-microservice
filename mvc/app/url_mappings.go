package app

import (
	"github.com/diogoramos/go-microservice/mvc/controllers"
)

func mapUrls(){
	router.GET("/users/:userId", controllers.GetUser)
}