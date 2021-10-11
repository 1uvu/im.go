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

// todo mod

func SaltPassword(passwordHash string) string {
	return passwordHash
}

func CheckPassword(saltedPassword string) bool {
	return saltedPassword == UnsaltPassword(saltedPassword)
}

func UnsaltPassword(saltedPassword string) string {
	return saltedPassword
}
