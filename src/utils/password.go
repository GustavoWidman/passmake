package utils

import (
	"crypto/rand"
	"math/big"
)

// GeneratePassword generates a secure random password of the specified length using the given charset.
func GeneratePassword(length uint8, charset string) string {
	b := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range b {
		num, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			// Handle the error appropriately in your application context.
			// For example: log the error and return an empty string.
			return ""
		}
		b[i] = charset[num.Int64()]
	}

	return string(b)
}
