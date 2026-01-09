package models

import (
	"database/sql"
	"errors"
	"time"
)

type Period struct {
	ID        string    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"user_id"`
	StartDate time.Time `db:"start_date" json:"start_date"`
	EndDate   time.Time `db:"end_date" json:"end_date"`
}

func GetCurrentPeriodByUserId(db *sql.DB, userId string) (Period, error) {

	//getting the latest period that the user is using
	query := `
		select * from periods where user_id = $1 order by end_date desc limit 1
	`

	period := Period{}
	err := db.QueryRow(query, userId).Scan(
		&period.ID,
		&period.UserID,
		&period.StartDate,
		&period.EndDate,
	)

	//handling errors
	if err == sql.ErrNoRows {
		return period, nil //no period found
	}

	if err != nil {
		return period, errors.New("error fetching period from database: " + err.Error())
	}

	return period, nil
}
