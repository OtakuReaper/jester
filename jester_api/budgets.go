package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jester/database"
	"jester/models"
)

type BudgetType struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type Budget struct {
	Id           string  `json:"id"`
	BudgetTypeId string  `json:"budget_type_id"`
	UserId       string  `json:"user_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Colour       string  `json:"colour"`
	Allocated    float64 `json:"allocated"`
	Spent        float64 `json:"spent"`
	Amount       float64 `json:"amount"`
}

func getBudgetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("I'm getting hit!")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//getting the param from the URL

	userId := r.PathValue("id")

	fmt.Println("User ID:", userId)

	if userId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	budgets := []models.Budget{}

	budgets, err := models.GetBudgetsByUserId(database.DB, userId)
	if err != nil {
		http.Error(w, "Error retrieving budgets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(budgets)
}
