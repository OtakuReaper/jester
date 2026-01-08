package main

import (
	"net/http"

	"jester/handlers"
)

func AddRoutes(mux *http.ServeMux) {

	mux.Handle("POST /auth/login", http.HandlerFunc(handlers.LoginHandler)) //TODO: add stronger auth handling
	mux.Handle("/auth/profile", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(handlers.ProfileHandler))))
	mux.Handle("GET /budgets/{id}", handlers.LoggingMiddleware(handlers.AuthMiddleware(http.HandlerFunc(getBudgetsHandler))))

	mux.Handle("/", http.HandlerFunc(handlers.HealthHandler))
	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))
	mux.Handle("/hello", http.HandlerFunc(handlers.HelloHandler))
	mux.Handle("/data", http.HandlerFunc(handlers.GettingDataHandler))
}
