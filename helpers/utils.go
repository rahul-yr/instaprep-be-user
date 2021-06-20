package helpers

import (
	"crypto/sha256"
	"fmt"
)

// Sha256 Algorithm
func Hash256(s string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
