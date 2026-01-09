package handlers

import (
	"encoding/json"
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
	entries, err := models.GetCurrentEntriesByUserId(database.DB, userId)
	if err != nil {
		http.Error(w, "Error retrieving entries: "+err.Error(), http.StatusInternalServerError) //TODO: remove the err.Error() in production for security reasons
		return
	}

	json.NewEncoder(w).Encode(entries)
}

func GetPeriodsHandler(w http.ResponseWriter, r *http.Request) {
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

	periods := []models.Period{}
	periods, err := models.GetPeriodsByUserId(database.DB, userId)
	if err != nil {
		http.Error(w, "Error retrieving periods: "+err.Error(), http.StatusInternalServerError) //TODO: remove the err.Error() in production for security reasons
		return
	}

	if len(periods) == 0 {
		json.NewEncoder(w).Encode([]models.Period{})
		return
	}

	//formatting periods for response
	type PeriodResponse struct {
		ID        string `json:"id"`
		Order     int    `json:"order"`
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date,omitempty"`
	}

	var formattedPeriods []PeriodResponse
	for index, period := range periods {
		formattedPeriod := PeriodResponse{
			ID:        period.ID,
			Order:     index + 1,
			StartDate: period.StartDate.Format("2006-01-02"),
			EndDate:   "",
		}

		if period.EndDate != nil && period.EndDate.Valid {
			formattedPeriod.EndDate = period.EndDate.Time.Format("2006-01-02")
		}

		formattedPeriods = append(formattedPeriods, formattedPeriod)
	}

	json.NewEncoder(w).Encode(formattedPeriods)
}
