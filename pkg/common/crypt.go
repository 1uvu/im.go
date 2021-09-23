package common

import (
	"crypto/sha1"
	"fmt"
)

func SHA1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	hashedText := h.Sum(nil)
	return fmt.Sprintf("%x", hashedText)
}

func SaltPassword(passwordHash string) string {
	return ""
}

func CheckPassword(saltedPassword string) bool {
	return false
}

func UnsaltPassword(saltedPassword string) string {
	return ""
}
