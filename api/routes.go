package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Holds All the Different API Routes and Route Setup Functions

const API_VERSION = "v1"

type LocalRouter struct {
	*mux.Router
}

// HandleFuncVersioned is a small wrapper function for Handle Func to append our API Version constant automagically
func (r *LocalRouter) HandleFuncVersioned(version, path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc("/"+version+path, f)
}

func setRoutes(router *LocalRouter) {
	//Setup Routes In Here
	//Example: SetContactUsRoutes()
}

/* Example:
func SetContactUsRoutes() {
	//Regex For If We Are Expecting an ID in the URL
	mysqlIdRegex := "(?:[0-9]{1,24})"

	expectedUrl := "/contact/")
	expectedUrlWithRegex := fmt.Sprintf("/contact/{contactID:%s}/", mysqlIdRegex)
	router.HandleFuncVersioned(
		API_VERSION,
		expectedUrl,
		contactHandler,).Methods("PUT", "GET").Name("contact")
	)
}

*/
