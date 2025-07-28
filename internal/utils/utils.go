package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func RandomFileName() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		n, err := rand.Int(rand.Reader, bigInt(len(letters)))
		if err != nil {
			panic(err)
		}
		b[i] = letters[n.Int64()]
	}
	return string(b)
}

func bigInt(n int) *big.Int {
	return big.NewInt(int64(n))
}

func FileSizeMB(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	sizeMB := float64(info.Size()) / (1024 * 1024)
	return fmt.Sprintf("%.2f MB", sizeMB), nil
}
