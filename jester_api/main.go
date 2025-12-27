package main

import (
	"log"
	"net/http"
	"time"
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

	// registering handlers with logging middleware
	mux.Handle("/", http.HandlerFunc(healthHandler))
	mux.Handle("/health", http.HandlerFunc(healthHandler))
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/data", http.HandlerFunc(gettingDataHandler))

	mux.Handle("/auth/login", http.HandlerFunc(loginHandler)) //TODO: add stronger auth handling
	mux.Handle("/auth/profile", loggingMiddleware(corsMiddleware(http.HandlerFunc(profileHandler))))

	log.Printf("Starting Jester API server on port %s...", port)

	// Start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
