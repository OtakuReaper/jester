package models

import (
	"database/sql"
	"errors"
	"time"
)

type Budget struct {
	ID            string         `db:"id" json:"id"`
	UserID        string         `db:"user_id" json:"user_id"`
	BudgetTypeID  string         `db:"budget_type_id" json:"budget_type_id"`
	Name          string         `db:"name" json:"name"`
	Description   string         `db:"description" json:"description"`
	Color         string         `db:"color" json:"color"`
	Allocation    float64        `db:"allocation" json:"allocation"`
	CurrentAmount float64        `db:"current_amount" json:"current_amount"`
	CreatedAt     time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time      `db:"updated_at" json:"updated_at"`
	CreatedBy     sql.NullString `db:"created_by" json:"created_by,omitempty"`
	UpdatedBy     sql.NullString `db:"updated_by" json:"updated_by,omitempty"`
}

func GetBudgetsByUserId(db *sql.DB, userID string) ([]Budget, error) {
	query := `
		select * from budgets where user_id = $1
	`

	budgets := []Budget{}
	rows, err := db.Query(query, userID)
	defer rows.Close()

	//reading the rows
	for rows.Next() {
		var budget Budget
		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.BudgetTypeID,
			&budget.Name,
			&budget.Description,
			&budget.Color,
			&budget.Allocation,
			&budget.CurrentAmount,
			budget.CreatedAt,
			&budget.UpdatedAt,
			&budget.CreatedBy,
			&budget.UpdatedBy,
		)
		if err != nil {
			return nil, errors.New("error scanning budget row: " + err.Error())
		}
		budgets = append(budgets, budget)
	}

	//handling errors
	if err == sql.ErrNoRows {
		return nil, nil //no budgets found
	}

	if err != nil {
		return nil, errors.New("error fetching budgets from database: " + err.Error())
	}

	return budgets, nil
}
