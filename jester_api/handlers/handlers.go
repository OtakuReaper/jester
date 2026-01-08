package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Middleware
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Handlers
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Jester API is healthy",
		Status:  "ok",
	}
	json.NewEncoder(w).Encode(response)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//ensuring that it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//preparing the response
	response := Response{
		Message: "Hello, welcome to Jester API!",
		Status:  "ok",
	}
	json.NewEncoder(w).Encode(response)
}

func GettingDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//ensuring that it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parsing the request body
	var requestData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	data := requestData["data"]
	if data == nil {
		http.Error(w, "'data' field is required", http.StatusBadRequest)
		return
	}

	log.Printf("Received data: %v", data)

	//preparing the response
	response := Response{
		Message: fmt.Sprintf("%v", data),
		Status:  "ok",
	}
	json.NewEncoder(w).Encode(response)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//ensuring that it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//checking for the auth_token cookie
	token, err := r.Cookie("auth_token")
	if err != nil || token.Value == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Println(token)

	//preparing the response
	response := map[string]interface{}{
		"id":        "someId", //TODO: replace with real user ID
		"username":  "admin",
		"status_id": "someId",            //TODO: replace with real status ID
		"email":     "admin@example.com", //TODO: replace with real email
	}
	json.NewEncoder(w).Encode(response)
}
