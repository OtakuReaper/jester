package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	budgets := []Budget{}

	budgets = append(budgets, Budget{
		Id:           "1",
		BudgetTypeId: "1",
		UserId:       "1",
		Name:         "Pool",
		Description:  "Where all money is kept, until allocation",
		Colour:       "#98ec94ff",
		Allocated:    271.23,
		Spent:        161.95,
		Amount:       109.28,
	})

	budgets = append(budgets, Budget{
		Id:           "2",
		BudgetTypeId: "2",
		UserId:       "1",
		Name:         "Land Debt",
		Description:  "For paying off the land loan",
		Colour:       "#ef8888ff",
		Allocated:    500.00,
		Spent:        500.00,
		Amount:       0.00,
	})

	//preparing the response

	json.NewEncoder(w).Encode(budgets)
}
