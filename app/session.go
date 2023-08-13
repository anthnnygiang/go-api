package app

//
//import (
//	"crypto/sha256"
//	"encoding/base64"
//	"github.com/rs/xid"
//)
//
//type Session struct {
//	ID        xid.ID
//	UserID    xid.ID
//	TokenHash string
//}
//
//func (s Session) CreateToken(token string) string {
//	tokenHash := sha256.Sum256([]byte(token))
//	return base64.URLEncoding.EncodeToString(tokenHash[:])
//}
//
//type SessionService interface {
//	CreateSession(UserID xid.ID) (*Session, string, error)
//	VerifySession(token string) (*User, error)
//	DeleteSession(token string) (*Session, error)
//}
