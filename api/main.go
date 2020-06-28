package main

import (
	"encoding/json"
	"flag"
	"go_server/config"
	"go_server/db"
	"go_server/log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

//spaHandler implements the http.Handler so we can use it
// to respond to HTTP requests. The path to the static directory and path
// to the index file within the static directory are used to
// serve the SPA in de given static directory

type spaHandler struct {
	staticPath string
	indexPath  string
}

/*
ServeHTTP inspects the URl path to locate a file within the static dir
on the SPA handler. If a file is found, it will be served. If not, the file
located at the indedx path on the SPA handler will be srved. This is suitable
behavior for serving an SPA
*/
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		//if we failed to get the absolute path respond with a 400 bad request and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a fle exists at a given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating
		// the file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config_path", "/config/server.toml", "Path to TOML Configuration File.")

	config.InitializeConfig(configPath)
	log.InitializeLog()

	db.CreateMySQLHandler(config.C.MySQL)

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	spa := spaHandler{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.C.Server.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

/*
package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gotilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// routes consist of a path and a handler function
	r.HandleFunc("/", YourHandler)

	// bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000",r ))
}

*/
