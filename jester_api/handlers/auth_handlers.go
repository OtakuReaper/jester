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
	log.Printf("Login attempt from IP: %v", ipAdress)

	//TODO: figure out how to handle input validation
	username := requestData["username"]
	password := requestData["password"]

	//checking on the database for the login request's account
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
		fmt.Println(err.Error())
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	//creating session
	newSession := models.NewSession{
		UserId:    user.ID,
		JwtToken:  "", //initially empty, will be set after JWT generation
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(6 * time.Minute), //TODO: change to a more appropriate time
	}

	sessionID, err := models.CreateSession(database.DB, newSession)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	token, err := security.GenerateJWT(*sessionID, user.ID)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	//updating session with the generated JWT
	err = models.UpdateSessionsToken(database.DB, *sessionID, token)
	if err != nil {
		http.Error(w, "Error updating session", http.StatusInternalServerError)
		return
	}

	//creating the cookie
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   6 * 60, //6 minutes //TODO: change to a more appropriate time
		Expires:  time.Now().Add(6 * time.Minute),
		HttpOnly: true,
		Secure:   false, //set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	//preparing the response
	//TODO send back the profile
	response := map[string]interface{}{
		"id":        user.ID,
		"username":  user.Username,
		"status_id": user.StatusID, //TODO: implement check for status in the future
		"email":     user.Email,
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

	//validate the token and check the session
	claims, err := security.ValidateJWT(token.Value)
	if err != nil {
		http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		return
	}

	session, err := models.GetSessionById(database.DB, claims.SessionId)
	if err != nil {
		http.Error(w, "Error fetching session", http.StatusInternalServerError)
		return
	}

	if session == nil {
		http.Error(w, "Session not found", http.StatusUnauthorized)
		return
	}

	//if the session is not expired return the user profile
	if session.ExpiresAt.Before(time.Now()) {
		http.Error(w, "Session expired", http.StatusUnauthorized)
		return
	}

	user, err := models.GetUserById(database.DB, session.UserId)
	if err != nil {
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	//preparing the response
	response := map[string]interface{}{
		"id":        user.ID,
		"username":  user.Username,
		"status_id": user.StatusID,
		"email":     user.Email,
	}
	json.NewEncoder(w).Encode(response)
}
