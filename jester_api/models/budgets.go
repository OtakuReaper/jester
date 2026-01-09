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
	PeriodID      sql.NullString `db:"period_id" json:"period_id,omitempty"`
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

type WorkingBudget struct {
	UserID        string         `db:"user_id" json:"user_id"`
	BudgetTypeID  string         `db:"budget_type_id" json:"budget_type_id"`
	PeriodID      sql.NullString `db:"period_id" json:"period_id,omitempty"`
	Name          string         `db:"name" json:"name"`
	Description   string         `db:"description" json:"description"`
	Color         string         `db:"color" json:"color"`
	Allocation    float64        `db:"allocation" json:"allocation"`
	CurrentAmount float64        `db:"current_amount" json:"current_amount"`
}

func CreateBudget(db *sql.DB, newBudget WorkingBudget) (Budget, error) {
	//inserting the new budget into the database
	query := `
		insert into budgets 
		(user_id, budget_type_id, period_id, name, description, color, allocation, current_amount)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		returning id, user_id, budget_type_id, period_id, name, description, color, allocation, current_amount, created_at, updated_at, created_by, updated_by
	`

	budget := Budget{}
	err := db.QueryRow(
		query,
		newBudget.UserID,
		newBudget.BudgetTypeID,
		newBudget.PeriodID,
		newBudget.Name,
		newBudget.Description,
		newBudget.Color,
		newBudget.Allocation,
		newBudget.CurrentAmount,
	).Scan(
		&budget.ID,
		&budget.UserID,
		&budget.BudgetTypeID,
		&budget.PeriodID,
		&budget.Name,
		&budget.Description,
		&budget.Color,
		&budget.Allocation,
		&budget.CurrentAmount,
		&budget.CreatedAt,
		&budget.UpdatedAt,
		&budget.CreatedBy,
		&budget.UpdatedBy,
	)

	if err != nil {
		return Budget{}, err
	}

	return budget, nil
}

// READ
func GetBudgetsByUserId(db *sql.DB, userID string) ([]Budget, error) {

	//getting the user's current period
	period, err := GetCurrentPeriodByUserId(db, userID)
	if err != nil {
		return nil, err
	}

	//if there is no current period, return empty slice
	if period.ID == "" {
		return []Budget{}, nil
	}

	query := `
		select * from budgets where user_id = $1
	`

	budgets := []Budget{}
	rows, err := db.Query(query, userID)
	defer rows.Close()

	//reading the rows
	for rows.Next() {
		budget := Budget{}

		err := rows.Scan(
			&budget.ID,
			&budget.UserID,
			&budget.BudgetTypeID,
			&budget.PeriodID,
			&budget.Name,
			&budget.Description,
			&budget.Color,
			&budget.Allocation,
			&budget.CurrentAmount,
			&budget.CreatedAt,
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

// UPDATE
func UpdateBudget(db *sql.DB, budgetID string, updatedBudget WorkingBudget) (Budget, error) {
	query := `
		update budgets set
		budget_type_id = coalesce($1, budget_type_id),
		period_id = coalesce($2, period_id),
		name = coalesce($3, name),
		description = coalesce($4, description),
		color = coalesce($5, color),
		allocation = coalesce($6, allocation),
		current_amount = coalesce($7, current_amount),
		updated_at = now()
		where id = $8
		returning id, user_id, budget_type_id, period_id, name, description, color, allocation, current_amount, created_at, updated_at, created_by, updated_by
	`

	budget := Budget{}
	err := db.QueryRow(
		query,
		updatedBudget.BudgetTypeID,
		updatedBudget.PeriodID,
		updatedBudget.Name,
		updatedBudget.Description,
		updatedBudget.Color,
		updatedBudget.Allocation,
		updatedBudget.CurrentAmount,
		budgetID,
	).Scan(
		&budget.ID,
		&budget.UserID,
		&budget.BudgetTypeID,
		&budget.PeriodID,
		&budget.Name,
		&budget.Description,
		&budget.Color,
		&budget.Allocation,
		&budget.CurrentAmount,
		&budget.CreatedAt,
		&budget.UpdatedAt,
		&budget.CreatedBy,
		&budget.UpdatedBy,
	)

	if err != nil {
		return Budget{}, errors.New("error updating budget: " + err.Error())
	}

	return budget, nil
}
