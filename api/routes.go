package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Holds All the Different API Routes and Route Setup Functions

const API_VERSION = "v1"

type LocalRouter struct {
	*mux.Router
}

func setRoutes(router *LocalRouter) {
	// http.Handle("/", router)

	router.HandleFunc("/", HomeHandler).Methods("GET", "PUT")
	router.HandleFunc("/contact", ContactHandler).Methods("GET", "PUT")

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Contact page!")

	// w.Header().Set("Content-Type", "application/json")
	//https://example.com/v2/contact?id=123
	/*
		json body: {
			name: someName,
			email: someEmail,
			message: someMessage
		}
	*/

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Welcome to home page!")
}
