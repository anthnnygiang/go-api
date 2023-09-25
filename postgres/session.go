package postgres

import (
	"anthnnygiang/api-template/app"
	"anthnnygiang/api-template/utils"
	"database/sql"
	"github.com/rs/xid"
)

type SessionService struct {
	DB *sql.DB
}

func (s *SessionService) CreateSession(userID string) (*app.Session, string, error) {
	id := xid.New().String()

	//Generate token and its hash
	rawToken, err := utils.CreateSessionToken()
	tokenHash := utils.HashToken(rawToken)
	session := app.Session{
		ID:        id,
		UserID:    userID,
		TokenHash: tokenHash,
	}

	//Insert into sessions table
	row := s.DB.QueryRow(`
		INSERT INTO sessions (id, user_id, token_hash)
		VALUES ($1, $2, $3) RETURNING token_hash;`, id, session.UserID, session.TokenHash)
	err = row.Scan(&tokenHash)
	if err != nil {
		return nil, "", err
	}
	if utils.HashToken(rawToken) != tokenHash {
		return nil, "", err
	}

	return &session, rawToken, nil
}
