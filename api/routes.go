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

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/contact", ContactHandler).Methods("GET", "PUT")

}

// // // defining autentication struct
// type authenticationMiddleware struct {
// 	tokenUsers map[string]string
// }

// //initializing
// func (amw *authenticationMiddleware) Populate() {
// 	amw.tokenUsers["00000000"] = "user0"
// 	amw.tokenUsers["aaaaaaaa"] = "userA"
// 	amw.tokenUsers["05f717e5"] = "randomUser"
// 	amw.tokenUsers["deadbeef"] = "user0"
// }

// // Middleware function which will be called for each request
// func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
// 	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("X-Session-Token")

// 		if user, found := amw.tokenUsers[token]; found {
// 			// we found the token in our map
// 			log.Printf("Authenticated user &s\n", user)
// 			// pass down the request to the next middleware (or final handler)
// 			next.ServeHTTP(w, r)
// 		} else {
// 			// write an error and stop the handler chain
// 			http.Error(w, "Forbidden", http.StatusForbidden)
// 		}
// 	})
// }

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		//Will Eventually Add Some Authentication Business

		var c models.Contact
		selectedContacts, err := c.SelectAllContacts()
		if err != nil {

			// serving HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			log.Errorf("HTTP Server Error Return 500: %v", err)
		}
		json.NewEncoder(w).Encode(selectedContacts)

	case "PUT":

		var c models.Contact

		id, err := c.PutContact()

		if err != nil {
			log.Errorf("Failed to insert into a contact into database")
		}

		log.Printf("Inserted row with ID of: %d\n", id)

	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// debuggin purposes
	fmt.Fprintln(w, "Welcome to the home page")

}
