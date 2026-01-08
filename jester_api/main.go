package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"jester/database"
	"jester/handlers"

	"github.com/joho/godotenv"
)

func main() {

	//Loading the Env variales from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//converting port from string to int
	databasePort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	// configuring database connection
	dbConfig := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     databasePort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}

	// connect to database
	if err := database.Connect(dbConfig); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//this will close the database connection when main exits
	defer database.Close()

	// creating a new ServeMux aka router
	mux := http.NewServeMux()

	port := "8080" //TODO: get from env variable
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handlers.CorsMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	AddRoutes(mux)

	// Start server
	log.Printf("Starting Jester API server on port %s...", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
