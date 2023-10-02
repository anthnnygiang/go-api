package postgres

import (
	"database/sql"
)

type SessionService struct {
	DB *sql.DB
}
