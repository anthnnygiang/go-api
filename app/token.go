package app

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"github.com/google/uuid"
	"time"
)

type Token struct {
	Hash   []byte
	UserID uuid.UUID
	Scope  Scope
	Expiry time.Time
}

type Scope string

const (
	ScopeActivate     Scope = "activate"
	ScopeAuthenticate Scope = "authenticate"
)

type TokenService interface {
	AddToken(token *Token) (*Token, error)
}

func GenerateToken(user *User, ttl time.Duration, scope Scope) (*Token, *string, error) {
	token := Token{
		UserID: user.ID,
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
	hash := sha256.Sum256([]byte(rawToken))
	token.Hash = hash[:]

	return &token, &rawToken, nil
}
