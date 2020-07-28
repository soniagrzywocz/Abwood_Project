package main

import (
	"encoding/json"
	"fmt"
	"go_server/db"
	"go_server/models"
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

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/contact", ContactHandler).Methods("GET", "PUT")

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintln(w, "Contact page!")

	w.Header().Set("Content-Type", "application/json")

	var c models.Contact
	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		fmt.Println("Error")
	}

	json.NewEncoder(w).Encode(c)

	// problem in passing here

	db.GetInfo(&c) // calling the mysql db

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to the home page")
	//w.Header().Set("Content-Type", "application/json")

	// probably not necessary here since it's just the home page, not the contact page

	// var c models.Contact
	// err := json.NewDecoder(r.Body).Decode(&c)

	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// json.NewEncoder(w).Encode(c)

}
