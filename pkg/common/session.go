package common

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const (
	SessionPrefix = "session_"
)

func CreateToken(length int) string {
	rd := make([]byte, length)
	io.ReadFull(rand.Reader, rd)
	return base64.URLEncoding.EncodeToString(rd)
}

func CreateSessionIDByToken(token string) string {
	return SessionPrefix + token
}

func CreateSessionIDByUserID(userID uint64) string {
	return fmt.Sprintf("%smap_%d", SessionPrefix, userID)
}

func GetSessionIDByToken(token string) string {
	return SessionPrefix + token
}
