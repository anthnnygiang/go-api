package postgres

import (
	"anthnnygiang/api-template/app"
	"database/sql"
	"fmt"
)

type UserService struct {
	DB *sql.DB
}

func (us *UserService) CreateUser(user *app.User) (*app.User, error) {

	//Insert into users table
	var u app.User
	row := us.DB.QueryRow(`
	  INSERT INTO users (id, created_at, email, password_hash, activated)
	  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, email, password_hash, activated`,
		user.ID, user.CreatedAt, user.Email, user.PasswordHash, user.Activated)

	err := row.Scan(&u.ID, &u.CreatedAt, &u.Email, &u.PasswordHash, &u.Activated)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &u, nil
}

func (us *UserService) ActivateUser(user *app.User) (*app.User, error) {
	return nil, nil
}
