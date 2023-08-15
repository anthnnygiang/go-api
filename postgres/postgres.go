package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("TESTDB_DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
