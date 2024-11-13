package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "developement", "Environment(developement|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: config,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %d", config.env, config.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
