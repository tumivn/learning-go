package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := app.route()

	addr := fmt.Sprintf(":%d", cfg.port)
	logger.Println("Starting %s server on port %s", cfg.env, addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		println(err)
	}
}

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}
