package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID           string         `db:"id" json:"id"`
	StatusID     string         `db:"status_id" json:"status_id"`
	Username     string         `db:"username" json:"username"`
	Email        string         `db:"email" json:"email"`
	PasswordHash string         `db:"password_hash" json:"-"`
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
		&user.Email,
		&user.PasswordHash,
		&user.OTPSecret,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.UpdatedBy,
	)

	fmt.Println(user.PasswordHash)

	//hanlding errors
	if err == sql.ErrNoRows {
		return nil, nil //user not found
	}

	if err != nil {
		return nil, errors.New("error fetching user from database: " + err.Error())
	}

	return user, nil
}

func GetUserById(db *sql.DB, userId string) (*User, error) {
	query := `
		select * from users where id = $1
	`

	user := &User{}
	err := db.QueryRow(query, userId).Scan(
		&user.ID,
		&user.StatusID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
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
