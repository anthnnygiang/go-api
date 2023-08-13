package postgres

import (
	"9z/go-api-template/app"
	"database/sql"
)

type UserService struct {
	DB *sql.DB
}

func (us *UserService) CreateUser(user *app.User) (*app.User, error) {
	newUser := app.User{
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
	//	Insert into db

	return &newUser, nil
}
