package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Session struct {
	Id        string    `db:"id" json:"id"`
	UserId    string    `db:"user_id" json:"user_id"`
	JwtToken  string    `db:"jwt_token" json:"token"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
}

type NewSession struct {
	UserId    string
	JwtToken  string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// CREATE
func CreateSession(db *sql.DB, newSess NewSession) (*string, error) {
	query := `
		insert into sessions (user_id, jwt_token, created_at, expires_at) values 
		($1, $2, $3, $4)
		returning id
	`

	sessionId := ""
	err := db.QueryRow(
		query,
		newSess.UserId,
		newSess.JwtToken,
		newSess.CreatedAt,
		newSess.ExpiresAt,
	).Scan(&sessionId)

	//handling errors
	if err != nil {
		return nil, errors.New("error creating session in database: " + err.Error())
	}

	return &sessionId, nil
}

// READ
func GetSessionById(db *sql.DB, sessionId string) (*Session, error) {
	query := `
		select * from sessions where id = $1
	`

	session := &Session{}
	err := db.QueryRow(query, sessionId).Scan(
		&session.Id,
		&session.UserId,
		&session.JwtToken,
		&session.CreatedAt,
		&session.ExpiresAt,
	)

	//handling errors
	if err == sql.ErrNoRows {
		return nil, nil //session not found
	}

	if err != nil {
		return nil, errors.New("error fetching session from database: " + err.Error())
	}

	return session, nil
}

// UPDATE
func UpdateSessionsToken(db *sql.DB, sessionId string, newToken string) error {
	query := `
		update sessions set jwt_token = $1 where id = $2
	`

	result, err := db.Exec(query, newToken, sessionId)

	//handling errors
	if err != nil {
		return fmt.Errorf("error updating session in database: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no session found with id: %s", sessionId)
	}

	return nil
}
