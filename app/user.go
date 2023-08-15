package app

type User struct {
	ID           string
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
