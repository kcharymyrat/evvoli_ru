package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"evvoli.ru/api"
)

func main() {
	fmt.Println("cmd/api - main")

	// Initialize configurations
	var cfg api.ApiConfig

	// Set the port num and environment for the config
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Set the logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create Dependency Injection for the api app
	app := &api.ApiApplication{
		Config: cfg,
		Logger: logger,
	}

	// New mux (multiplexer, router), and handlers (view)
	mux := app.Routes()

	// Declare the server with the appropriate settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server
	logger.Printf("starting %s server on %s", cfg.Env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
