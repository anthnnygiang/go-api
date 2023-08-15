package app

type Session struct {
	ID        string
	UserID    string
	TokenHash string
}

type SessionService interface {
	CreateSession(userID string) (*Session, string, error)
	//VerifySession(token string) (string, error)
	//DeleteSession(token string) (*Session, error)
}
