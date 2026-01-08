package models

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID           string         `db:"id" json:"id"`
	StatusID     string         `db:"status_id" json:"status_id"`
	Username     string         `db:"username" json:"username"`
	PasswordHash string         `db:"password_hash" json:"-"`
	Email        string         `db:"email" json:"email"`
	OTPSecret    sql.NullString `db:"otp_secret" json:"-,omitempty"`
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at" json:"updated_at"`
	CreatedBy    sql.NullString `db:"created_by" json:"created_by,omitempty"`
	UpdatedBy    sql.NullString `db:"updated_by" json:"updated_by,omitempty"`
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := `
		select * from users where username = $1
	`

	user := &User{}
	err := db.QueryRow(query, username).Scan(
		&user.ID,
		&user.StatusID,
		&user.Username,
		&user.PasswordHash,
		&user.Email,
		&user.OTPSecret,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
	)

	//hanlding errors
	if err == sql.ErrNoRows {
		return nil, nil //user not found
	}

	if err != nil {
		return nil, errors.New("error fetching user from database: " + err.Error())
	}

	return user, nil
}
