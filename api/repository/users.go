package repository

import (
	"context"
	"database/sql"
	"time"
	"log"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	Id string `db:"id" json:"id"`
	RoleId string `db:"role_id" json:"role_id"`
	Username string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"created_by"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type UserRepo struct {
	db *sql.DB
	psql sq.StatementBuilderType
}

//implement functions here ...

func (r *UserRepo) GetAllUsers (ctx context.Context) ([]User, error){

	rows, err := r.psql.
		Select(
			"id", "role_id", "username", 
			"email", "created_at", "updated_at", 
			"created_by", "updated_by").
		From("users").RunWith(r.db).QueryContext(ctx)

	if err != nil {
		log.Println("Error querying users:", err)
		return nil, err
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Id, &user.RoleId, &user.Username,
			&user.Email, &user.CreatedAt, &user.UpdatedAt,
			&user.CreatedBy, &user.UpdatedBy,
		); err != nil {
			log.Println("Error scanning user row:", err)
			return nil, err
		}
		result = append(result, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Row iteration error:", err)
		return nil, err
	}
	
	return result, nil
}