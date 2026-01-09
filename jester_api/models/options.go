package models

import (
	"database/sql"
	"errors"
)

type EntryType struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type BudgetType struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

// READ
func GetEntryTypes(db *sql.DB) ([]EntryType, error) {
	query := `
		select * from entry_types
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New("error fetching entry types from database: " + err.Error())
	}
	defer rows.Close()

	entryTypes := []EntryType{}

	for rows.Next() {
		var et EntryType
		err := rows.Scan(
			&et.ID,
			&et.Name,
			&et.Description,
		)
		if err != nil {
			return nil, errors.New("error scanning entry type row: " + err.Error())
		}
		entryTypes = append(entryTypes, et)
	}

	return entryTypes, nil
}

func GetBudgetTypes(db *sql.DB) ([]BudgetType, error) {
	query := `
		select * from budget_types
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New("error fetching budget types from database: " + err.Error())
	}
	defer rows.Close()

	budgetTypes := []BudgetType{}

	for rows.Next() {
		var bt BudgetType
		err := rows.Scan(
			&bt.ID,
			&bt.Name,
			&bt.Description,
		)
		if err != nil {
			return nil, errors.New("error scanning budget type row: " + err.Error())
		}
		budgetTypes = append(budgetTypes, bt)
	}

	return budgetTypes, nil
}
