package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"jester/database"
	"jester/models"
	"jester/security"
)

// middleware and handlers for authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//trying to get the auth_token cookie
		token, err := r.Cookie("auth_token")
		if err != nil || token.Value == "" {
			fmt.Println("Token is not provided!")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

	//parsing the ip address from the request
	ipAdress := r.RemoteAddr

	//some logic to create a session

	log.Printf("Login attempt from IP: %v", ipAdress)

	//TODO: figure out how to handle input validation
	username := requestData["username"]
	password := requestData["password"]

	//checking on the database for the login request's account

	log.Printf("Login attempt for user: %v", username)
	log.Printf("Password provided: %v", password) //remove this in production!

	//TODO: some logic to verify username and password on the database

	newUser := models.User{}
	user := &newUser

	//fetching user from database
	user, err := models.GetUserByUsername(database.DB, username.(string))
	if err != nil {
		http.Error(w, "Error fetching user from database", http.StatusInternalServerError)
		return
	}

	//checking the password
	err = security.VerifyPassword(user.PasswordHash, password.(string))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	fmt.Println("Fetched User from Database")

	//TODO: some logic to create a session on the database

	//creating a JWT token (dummy token for now) for the response's cookie and body
	token := "dummy-jwt-token"

	//creating the cookie
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   6 * 60, //6 minutes
		Expires:  time.Now().Add(6 * time.Minute),
		HttpOnly: true,
		Secure:   false, //set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	//preparing the response
	//TODO send back the profile
	response := map[string]interface{}{
		"id":        "someId", //TODO: replace with real user ID
		"username":  "admin",
		"status_id": "someId",            //TODO: replace with real status ID
		"email":     "admin@example.com", //TODO: replace with real email
	}

	json.NewEncoder(w).Encode(response)
}
