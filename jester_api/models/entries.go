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

type WorkingEntry struct {
	BudgetID     string    `db:"budget_id" json:"budget_id"`
	BudgetTypeID string    `db:"budget_type_id" json:"budget_type_id"`
	PeriodID     string    `db:"period_id" json:"period_id"`
	Description  string    `db:"description" json:"description"`
	Date         time.Time `db:"date" json:"date"`
	Amount       float64   `db:"amount" json:"amount"`
}

// CREATE
func CreateEntry(db *sql.DB, newEntry WorkingEntry) (Entry, error) {
	//inserting the new entry into the database
	query := `
		insert into entries (budget_id, budget_type_id, period_id, description, date, amount) values
		($1, $2, $3, $4, $5, $6)
		returning id, budget_id, budget_type_id, period_id, description, date, amount
	`

	entry := Entry{}
	err := db.QueryRow(
		query,
		newEntry.BudgetID,
		newEntry.BudgetTypeID,
		newEntry.PeriodID,
		newEntry.Description,
		newEntry.Date,
		newEntry.Amount,
	).Scan(
		&entry.ID,
		&entry.BudgetID,
		&entry.BudgetTypeID,
		&entry.PeriodID,
		&entry.Description,
		&entry.Date,
		&entry.Amount,
	)

	//handling errors
	if err != nil {
		return Entry{}, errors.New("error creating entry in database: " + err.Error())
	}

	return entry, nil
}

// READ
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

// UPDATE
func UpdateEntry(db *sql.DB, entryId string, updatedEntry WorkingEntry) (Entry, error) {
	//updating the entry in the database
	query := `
		update entries
		set budget_id = coalesce($1, budget_id),
		    budget_type_id = coalesce($2, budget_type_id),
		    period_id = coalesce($3, period_id),
		    description = coalesce($4, description),
		    date = coalesce($5, date),
		    amount = coalesce($6, amount)
		where id = $7
		returning id, budget_id, budget_type_id, period_id, description, date, amount
	`

	entry := Entry{}
	err := db.QueryRow(
		query,
		updatedEntry.BudgetID,
		updatedEntry.BudgetTypeID,
		updatedEntry.PeriodID,
		updatedEntry.Description,
		updatedEntry.Date,
		updatedEntry.Amount,
		entryId,
	).Scan(
		&entry.ID,
		&entry.BudgetID,
		&entry.BudgetTypeID,
		&entry.PeriodID,
		&entry.Description,
		&entry.Date,
		&entry.Amount,
	)

	//handling errors
	if err != nil {
		return Entry{}, errors.New("error updating entry in database: " + err.Error())
	}

	return entry, nil
}

// DELETE
func DeleteEntryById(db *sql.DB, entryId string) error {
	//deleting the entry from the database
	query := `
		delete from entries where id = $1
	`

	_, err := db.Exec(query, entryId)

	//handling errors
	if err != nil {
		return errors.New("error deleting entry from database: " + err.Error())
	}

	return nil
}
