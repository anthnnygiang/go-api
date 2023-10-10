package postgres

import (
	"anthnnygiang/api-template/app"
	"database/sql"
)

type TokenService struct {
	DB *sql.DB
}

func (ts TokenService) AddToken(token *app.Token) (*app.Token, error) {
	return nil, nil
}
