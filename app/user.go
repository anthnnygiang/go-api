package app

import "github.com/rs/xid"

type User struct {
	ID           xid.ID
	Email        string
	PasswordHash string
}

type UserService interface {
	CreateUser(user *User) (*User, error)
	//UpdateUser(user *User) (*User, error)
	//DeleteUser(user *User) (*User, error)
	//AuthenticateUser(user *User) (*User, error)
	//UpdateUserPassword(user *User, newPassword string) (*User, error)
}
