package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type UserRole struct {
	Id          string `db:"id" json:"id"`
	RoleName    string `db:"role_name" json:"role_name"`
	Description string `db:"description" json:"description"`
}

type UserRoleRepo struct {
	db *sql.DB
	psql sq.StatementBuilderType
}

//implement functions here ...