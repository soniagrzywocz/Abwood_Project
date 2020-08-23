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
	router.HandleFunc("/inventory", InventoryHandler).Methods("GET", "PUT", "PATCH", "DELETE")

}

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
			log.Errorf("Failed to insert a contact into database")
		}

		log.Printf("Inserted row with ID of: %d\n", id)

	}

}

func InventoryHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Welcome to the inventory page") //works

	w.Header().Set("Content-Type", "application/json")

	// // methods needed - GET, PUT, PATCH, DELETE
	// // think about GET for all and GET for one

	switch r.Method {
	case "GET":
		// to do later here: add some authentication business
		var i models.InventoryItem
		selectedInventoryItems, err := i.SelectAllInventory()
		if err != nil {

			// serving the HTTP 500 error
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Errorf("HTTP Server Error return 500 for Inventory method GET: %v", err)
		}

		json.NewEncoder(w).Encode(selectedInventoryItems)

	case "PUT":
		// same logic as PUT for contact but doesn't work. Check for problems here
		var i models.InventoryItem
		err := i.PutInventoryItem()

		if err != nil {
			log.Errorf("Failed to Insert inventory item into database")
		}

		log.Printf("Inserted inventoryItem with ID of: %d\n", i.ID)

	case "PATCH":
		var i models.InventoryItem

		json.NewDecoder(r.Body).Decode(&i) //reading the updates heres
		updatedItem, err := i.UpdateItem(&i)

		if err != nil {
			log.Errorf("Error updating an existing element.")
		}

		// json.NewEncoder(w).Encode(updatedItem)
		log.Printf("An item updated in the database with an ID of: %d\n", updatedItem)

	case "DELETE":
		var i models.InventoryItem
		json.NewDecoder(r.Body).Decode(&i)
		itemToDelete := i.ID
		deleteItem, err := i.DeleteItem(itemToDelete)

		if err != nil {
			log.Errorf("Error deleting an item from a database")
		} else {
			log.Printf("An item deleted from a database with an ID of: %d\n", deleteItem)
		}

	} // end of switch

} // end of InventoryHandler

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// debuggin purposes
	fmt.Fprintln(w, "Welcome to the home page")

}
