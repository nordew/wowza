package generator

import (
	"crypto/rand"
	"fmt"
	"io"
)

type CharType int

const (
	NumbersOnly CharType = iota
	LettersOnly
	Mixed
)

var (
	numbers = []byte("0123456789")
	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	mixed   = append(numbers, letters...)
)

func GenerateCode(size int, charType CharType) (string, error) {
	var table []byte
	switch charType {
	case NumbersOnly:
		table = numbers
	case LettersOnly:
		table = letters
	case Mixed:
		table = mixed
	default:
		return "", fmt.Errorf("generator: invalid character type")
	}

	b := make([]byte, size)
	n, err := io.ReadAtLeast(rand.Reader, b, size)
	if n != size {
		return "", fmt.Errorf("generator: could not read enough random bytes: %w", err)
	}

	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b), nil
} 