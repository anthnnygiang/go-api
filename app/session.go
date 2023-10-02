package app

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"github.com/google/uuid"
	"time"
)

type Session struct {
	RawToken  string
	TokenHash [32]byte
	UserID    uuid.UUID
	Expiry    time.Time
	Scope     string
}

type SessionService interface {
	//CreateSession(userID string) (session *Session, rawToken string, err error)
	//VerifySession(rawToken string) (*User, error)
	//DeleteSession(rawToken string) (*Session, error)
}

func GenerateToken(user *User, ttl time.Duration, scope string) (*Session, error) {
	token := Session{
		UserID: user.ID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	// Initialize a zero-valued byte slice with a length of 16 bytes
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	//Set the plaintext and its hash
	token.RawToken = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	token.TokenHash = sha256.Sum256([]byte(token.RawToken))

	return &token, nil
}
