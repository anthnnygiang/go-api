package postgres

import (
	"anthnnygiang/api-template/app"
	"database/sql"
	"fmt"
)

type TokenService struct {
	DB *sql.DB
}

func (ts TokenService) AddToken(token *app.Token) (*app.Token, error) {
	var t app.Token
	row := ts.DB.QueryRow(`
	  INSERT INTO tokens (token_hash, user_id, scope, expiry)
	  VALUES ($1, $2, $3, $4 ) RETURNING token_hash, user_id, scope, expiry`,
		token.Hash, token.UserID, token.Scope, token.Expiry)

	err := row.Scan(&t.Hash, &t.UserID, &t.Scope, &t.Expiry)
	if err != nil {
		return nil, fmt.Errorf("create token: %w", err)
	}
	return &t, nil
}
