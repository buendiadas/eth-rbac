package main

import (
	"eth-rbac/models"
	"eth-rbac/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	models.Init()

	r := mux.NewRouter()
	r.HandleFunc("/signin", routes.Signin).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},   // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
	})

	err := http.ListenAndServe(":8000", c.Handler(r)) //Launch the app, visit localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
