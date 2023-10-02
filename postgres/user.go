package postgres

import (
	"anthnnygiang/api-template/app"
	"database/sql"
	"fmt"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) CreateUser(user *app.User) (*app.User, error) {

	//Insert into users table
	//Add an "activated "column
	var u app.User
	row := s.DB.QueryRow(`
	  INSERT INTO users (id, email, password_hash)
	  VALUES ($1, $2, $3) RETURNING id, email, password_hash`, user.ID, user.Email, user.PasswordHash)
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &u, nil
}
