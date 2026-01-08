package main

import (
	"net/http"

	h "jester/handlers"
)

func AddRoutes(mux *http.ServeMux) {

	mux.Handle("POST /auth/login", http.HandlerFunc(h.LoginHandler)) //TODO: add stronger auth handling
	mux.Handle("/auth/profile", h.LoggingMiddleware(h.AuthMiddleware(http.HandlerFunc(h.ProfileHandler))))
	mux.Handle("GET /budgets/{id}", h.LoggingMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetBudgetsHandler))))
	mux.Handle("GET /entries/{id}", h.LoggingMiddleware(h.AuthMiddleware(http.HandlerFunc(h.GetEntriesHandler))))

	mux.Handle("/", http.HandlerFunc(h.HealthHandler))
	mux.Handle("/health", http.HandlerFunc(h.HealthHandler))
	mux.Handle("/hello", http.HandlerFunc(h.HelloHandler))
	mux.Handle("/data", http.HandlerFunc(h.GettingDataHandler))
}
