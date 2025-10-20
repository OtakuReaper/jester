package database

import (
	"context"
	"database/sql"
	"fmt"
	// "os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDatabase(ctx context.Context) (*sql.DB, error) {
	// connectionString := os.Getenv("DATABASE_URL")
	connectionString := "postgresql://[user]:[password]@localhost:8000/jester_db?sslmode=disable"
	if connectionString == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(30 * time.Minute)

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctxPing); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	return db, nil
}