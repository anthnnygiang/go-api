package app

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

type Token struct {
	Plaintext string
	Hash      [32]byte
	UserID    string
	Expiry    time.Time
	Scope     string
}

func GenerateToken(user *User, ttl time.Duration, scope string) (*Token, error) {
	token := Token{
		UserID: user.ID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	// Initialize a zero-valued byte slice with a length of 32 bytes
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	//Set the plaintext and its hash
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	token.Hash = sha256.Sum256([]byte(token.Plaintext))

	return &token, nil
}
