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

// HandleFuncVersioned is a small wrapper function for Handle Func to append our API Version constant automagically
func (r *LocalRouter) HandleFuncVersioned(version, path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.HandleFunc("/"+version+path, f)
}

func setRoutes(router *LocalRouter) {

	r := mux.NewRouter()
	route := r.NewRoute().HeadersRegexp("Origin", "^https://apprenticenship.com$")

	yes, _ := http.NewRequest("GET", "apprenticenship.com", nil)
	yes.Header.Set("Origin", "https://apprenticenship.com")

	matchInfo := &mux.RouteMatch{}

	fmt.Printf("Match: %v %q\n", route.Match(yes, matchInfo), yes.Header["Origin"])

	//Setup Routes In Here
	//Example: SetContactUsRoutes()
}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	//vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	name := query.Get("name")
	email := query.Get("email")
	message := query.Get("message")

	w.Write([]byte(fmt.Sprintf(`{"name": %s, "email": %s, "message": %s}`, name, email, message)))

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
