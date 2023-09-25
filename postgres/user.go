package postgres

import (
	"anthnnygiang/api-template/app"
	"database/sql"
	"fmt"
	"github.com/rs/xid"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) CreateUser(user *app.User) (*app.User, error) {
	id := xid.New().String()

	//Insert into users table
	var userID string
	row := s.DB.QueryRow(`
	  INSERT INTO users (id, email, password_hash)
	  VALUES ($1, $2, $3) RETURNING id`, id, user.Email, user.PasswordHash)
	err := row.Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	u := app.User{ID: userID, Email: user.Email, PasswordHash: user.PasswordHash}
	return &u, nil
}
