package app

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"github.com/google/uuid"
	"time"
)

type Token struct {
	TokenHash [32]byte
	userID    uuid.UUID
	Expiry    time.Time
	Scope     string
}

func GenerateToken(user *User, ttl time.Duration, scope string) (*Token, *string, error) {
	token := Token{
		userID: user.ID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	// Initialize a zero-valued byte slice with a length of 16 bytes
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, nil, err
	}

	//Set the plaintext and its hash
	rawToken := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	token.TokenHash = sha256.Sum256([]byte(rawToken))

	return &token, &rawToken, nil
}
