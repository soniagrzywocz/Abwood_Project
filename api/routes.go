package main

import (
	"encoding/json"
	"fmt"
	"go_server/log"
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
	switch r.Method {
	case "GET":
		//Will Eventually Add Some Authentication Business
		var c models.Contact
		selectedContacts, err := c.SelectAllContacts()
		if err != nil {
			//Write some http return code here usually gonna be some form of
			//500 in this case as it means we failed to go to the db
			log.Errorf("HTTP Server Error Return 500: %v", err)
		}
		json.NewEncoder(w).Encode(selectedContacts)
	}
	// var c models.Contact
	// err := json.NewDecoder(r.Body).Decode(&c)

	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// problem in passing here

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
