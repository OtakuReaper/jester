package main

import (
	"log"
	"net/http"
	"time"

	"jester/handlers"
)

func main() {

	// creating a new ServeMux aka router
	mux := http.NewServeMux()

	port := "8080"
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	mux.Handle("POST /auth/login", http.HandlerFunc(handlers.LoginHandler)) //TODO: add stronger auth handling
	mux.Handle("/auth/profile",
		loggingMiddleware(
			handlers.AuthMiddleware(
				http.HandlerFunc(
					profileHandler))))

	mux.Handle("GET /budgets/{id}",
		loggingMiddleware(
			handlers.AuthMiddleware(
				http.HandlerFunc(
					getBudgetsHandler))))

	mux.Handle("/", http.HandlerFunc(healthHandler))
	mux.Handle("/health", http.HandlerFunc(healthHandler))
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/data", http.HandlerFunc(gettingDataHandler))

	log.Printf("Starting Jester API server on port %s...", port)

	// Start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
