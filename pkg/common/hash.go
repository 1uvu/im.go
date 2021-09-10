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

func CityHash32(bs []byte, bl uint32) uint32 {
	// todo 1 替换算法, 使用到 skiplist 和 rehash
	return 0
}
