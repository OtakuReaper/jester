package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jester/database"
	"jester/models"
)

func GetBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//only allowing GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//getting the param from the URL
	userId := r.PathValue("id")
	if userId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	budgets := []models.Budget{}

	budgets, err := models.GetBudgetsByUserId(database.DB, userId)
	if err != nil {
		http.Error(w, "Error retrieving budgets: "+err.Error(), http.StatusInternalServerError) //TODO: remove the err.Error() in production for security reasons
		return
	}

	json.NewEncoder(w).Encode(budgets)
}

func GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Getting hit!")

	//only allowing GET method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//getting the param from the URL
	userId := r.PathValue("id")
	if userId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	entries := []models.Entry{}

	fmt.Println("Failing Here!")

	entries, err := models.GetCurrentEntriesByUserId(database.DB, userId)
	if err != nil {
		http.Error(w, "Error retrieving entries: "+err.Error(), http.StatusInternalServerError) //TODO: remove the err.Error() in production for security reasons
		return
	}

	json.NewEncoder(w).Encode(entries)
}
