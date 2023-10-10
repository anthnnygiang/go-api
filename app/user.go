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
	AddUser(user *User) (*User, error)
	//ActivateUser(user *User) (*User, error)
	//ResetUserPassword(user *User) (*User, error)
	//UpdateUser(user *User) (*User, error)
	//DeleteUser(user *User) (*User, error)
	//AuthenticateUser(user *User) (*User, error)
}
