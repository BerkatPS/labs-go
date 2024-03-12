package main

import (
	"github.com/BerkatPS/go-jwt-testing/controllers/authcontroller"
	"github.com/BerkatPS/go-jwt-testing/controllers/dashboardcontroller"
	"github.com/BerkatPS/go-jwt-testing/middleware"
	"github.com/BerkatPS/go-jwt-testing/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	models.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	// membuat middleware untuk route home
	// mirip laravel , bisa menggunakan sub route
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/home", dashboardcontroller.Index).Methods("GET")
	api.Use(middleware.JWTMiddleware) // ini untuk menggunakan middleware
	log.Fatal(http.ListenAndServe(":8080", r))

}
