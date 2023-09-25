package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func CreateSessionToken() (string, error) {
	//number of bytes used to generate the token
	n := 32
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	if nRead < n {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func HashToken(rawToken string) string {
	tokenHash := sha256.Sum256([]byte(rawToken))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
