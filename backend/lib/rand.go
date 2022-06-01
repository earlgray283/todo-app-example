package lib

import (
	"crypto/rand"
	"math/big"
)

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		randi, _ := rand.Int(rand.Reader, big.NewInt(int64(len(LETTERS))))
		b[i] = LETTERS[randi.Int64()]
	}
	return b
}
