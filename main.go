package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rmsubekti/ghapis/app"
	"github.com/rmsubekti/ghapis/app/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/user/new", controllers.SignUp).Methods("POST")
	router.HandleFunc("/v1/user/me", controllers.Me).Methods("GET")
	router.HandleFunc("/v1/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/v1/product/new", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/v1/product/all", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/v1/product/{id}", controllers.GetProduct).Methods("GET")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
