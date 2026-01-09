package models

import (
	"database/sql"
	"errors"
	"time"
)

type Entry struct {
	ID           string    `db:"id" json:"id"`
	BudgetID     string    `db:"budget_id" json:"budget_id"`
	BudgetTypeID string    `db:"budget_type_id" json:"budget_type_id"`
	PeriodID     string    `db:"period_id" json:"period_id"`
	Description  string    `db:"description" json:"description"`
	Date         time.Time `db:"date" json:"date"`
	Amount       float64   `db:"amount" json:"amount"`
}

func GetCurrentEntriesByUserId(db *sql.DB, userId string) ([]Entry, error) {

	//getting the user's current period
	period, err := GetCurrentPeriodByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	//checking if there even is a current period
	if period.ID == "" {
		return []Entry{}, nil //no current period, return empty slice
	}

	//getting the entries for the user's current period
	query := `
		select * from entries where period_id = $1
	`

	entries := []Entry{}
	rows, err := db.Query(query, period.ID)
	defer rows.Close()

	//reading the rows
	for rows.Next() {
		entry := Entry{}

		err := rows.Scan(
			&entry.ID,
			&entry.BudgetID,
			&entry.BudgetTypeID,
			&entry.PeriodID,
			&entry.Description,
			&entry.Date,
			&entry.Amount,
		)
		if err != nil {
			return nil, errors.New("error scanning entry row: " + err.Error())
		}
		entries = append(entries, entry)
	}

	//handling errors
	if err == sql.ErrNoRows {
		return entries, nil //no entries found
	}

	if err != nil {
		return nil, errors.New("error fetching entries from database: " + err.Error())
	}

	return entries, nil
}
