package app

type Session struct {
	ID        string
	UserID    string
	TokenHash string
}

type SessionService interface {
	CreateSession(userID string) (session *Session, rawToken string, err error)
	//VerifySession(rawToken string) (*User, error)
	//DeleteSession(rawToken string) (*Session, error)
}
