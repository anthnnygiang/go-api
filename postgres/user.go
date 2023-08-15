package postgres

import (
	"9z/go-api-template/app"
	"database/sql"
	"fmt"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) CreateUser(user *app.User) (*app.User, error) {
	u := app.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}

	//Insert into users table
	var userID string
	row := s.DB.QueryRow(`
	  INSERT INTO users (id, email, password_hash)
	  VALUES ($1, $2, $3) RETURNING id`, u.ID, u.Email, u.PasswordHash)
	err := row.Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &u, nil
}
