package app

import (
	"github.com/diogoramos/go-microservice/mvc/controllers"
	"net/http"
)

func StartApp(){
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil ); err != nil {
		panic(err)
	}
}
