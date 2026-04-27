// Package util предоставляет дополнительные функции.
package util

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString генерирует случайную строку заданной длины
func GenerateRandomString(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("n must be greater than 0")
	}

	result := make([]byte, n)
	for i := range result {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}
	return string(result), nil
}
