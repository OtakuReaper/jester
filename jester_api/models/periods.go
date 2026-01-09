package models

import (
	"database/sql"
	"errors"
	"time"
)

type Period struct {
	ID        string        `db:"id" json:"id"`
	UserID    string        `db:"user_id" json:"user_id"`
	StartDate time.Time     `db:"start_date" json:"start_date"`
	EndDate   *sql.NullTime `db:"end_date" json:"end_date,omitempty"`
}

type NewPeriod struct {
	UserID    string    `db:"user_id" json:"user_id"`
	StartDate time.Time `db:"start_date" json:"start_date"`
}

type UpdatePeriod struct {
	StartDate *time.Time    `db:"start_date" json:"start_date,omitempty"`
	EndDate   *sql.NullTime `db:"end_date" json:"end_date,omitempty"`
}

// CREATE
func CreatePeriod(db *sql.DB, newPeriod NewPeriod) (Period, error) {
	//inserting the new period into the database
	query := `
		insert into periods (user_id, start_date)
		values ($1, $2)
		returning id, user_id, start_date, end_date
	`

	period := Period{}
	err := db.QueryRow(
		query,
		newPeriod.UserID,
		newPeriod.StartDate,
	).Scan(
		&period.ID,
		&period.UserID,
		&period.StartDate,
		&period.EndDate,
	)

	//handling errors
	if err != nil {
		return Period{}, errors.New("error creating period in database: " + err.Error())
	}

	return period, nil
}

// READ
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

func GetPeriodsByUserId(db *sql.DB, userId string) ([]Period, error) {

	//getting all periods for the user
	query := `
		select * from periods where user_id = $1 order by start_date desc
	`
	periods := []Period{}
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, errors.New("error fetching periods from database: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var period Period
		err := rows.Scan(
			&period.ID,
			&period.UserID,
			&period.StartDate,
			&period.EndDate,
		)
		if err != nil {
			return nil, errors.New("error scanning period from database: " + err.Error())
		}
		periods = append(periods, period)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error iterating over periods from database: " + err.Error())
	}

	return periods, nil
}

// UPDATE
func UpdatePeriodById(db *sql.DB, periodId string, updatedPeriod UpdatePeriod) (Period, error) {
	//updating the period in the database
	query := `
		update periods
		set start_date = coalesce($1, start_date),
		    end_date = coalesce($2, end_date)
		where id = $3
		returning id, user_id, start_date, end_date
	`

	period := Period{}
	err := db.QueryRow(
		query,
		updatedPeriod.StartDate,
		updatedPeriod.EndDate,
		periodId,
	).Scan(
		&period.ID,
		&period.UserID,
		&period.StartDate,
		&period.EndDate,
	)

	//handling errors
	if err != nil {
		return Period{}, errors.New("error updating period in database: " + err.Error())
	}

	return period, nil
}
