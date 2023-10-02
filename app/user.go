package app

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	Email        string
	PasswordHash []byte
	Activated    bool
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	//UpdateUser(user *User) (*User, error)
	//DeleteUser(user *User) (*User, error)
	//AuthenticateUser(user *User) (*User, error)
	//UpdateUserPassword(user *User, newPassword string) (*User, error)
}
