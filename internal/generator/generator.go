package generator

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword(length int, upper, lower, digits, symbols bool) string {
	var charset string

	if upper {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if lower {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if digits {
		charset += "0123456789"
	}
	if symbols {
		charset += "!@#$%^&*()-_=+[]{}<>?/|"
	}

	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	result := make([]byte, length)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[n.Int64()]
	}
	return string(result)
}
