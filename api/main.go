package main

import (
	"context"
	"flag"
	"go_server/config"
	"go_server/db"
	"go_server/log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	var configPath string
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "config_path", "", "Path to TOML Configuration File.")
	flag.Parse()
	config.InitializeConfig(configPath)
	log.InitializeLog()

	db.CreateMySQLHandler(config.C.MySQL)

	router := &LocalRouter{mux.NewRouter()}
	setRoutes(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.C.Server.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	//Graceful shutdown handling with ctrl+c and the various signal kill calls
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	//Block Until Signal Is Recieved
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
