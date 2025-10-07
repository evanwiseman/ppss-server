package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/evanwiseman/ppss-server/internal/config"
	"github.com/evanwiseman/ppss-server/internal/database"
	"github.com/evanwiseman/ppss-server/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const ()

func main() {
	// Load .env once
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using OS env variables")
	}

	// Load static configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Open DB once
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Wrap DB with queries
	queries := database.New(db)

	// Create local server with shared queries
	localSrv, err := server.NewLocalServer(cfg, queries)
	if err != nil {
		log.Fatal(err)
	}
	localMux := http.NewServeMux()
	server.LocalRoutes(localSrv, localMux)

	// Create public server with shared queries
	publicSrv, err := server.NewPublicServer(cfg, queries)
	if err != nil {
		log.Fatal(err)
	}
	publicMux := http.NewServeMux()
	server.PublicRoutes(publicSrv, publicMux)

	// Channel to log errors
	errs := make(chan error, 2)

	// Goroutine for local server
	go func() {
		log.Printf("Local server on %s", cfg.LocalAddress)
		errs <- http.ListenAndServe(cfg.LocalAddress, localMux)
	}()

	// Goroutine for public server
	go func() {
		log.Printf("Public server on %s", cfg.PublicAddress)
		errs <- http.ListenAndServe(cfg.PublicAddress, publicMux)
	}()

	log.Fatal(<-errs) // block until one server fails
}
