package app

import (
	"github.com/KatherineEbel/go-microservices/mvc/controllers"
	"net/http"
)

func Start() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
