package onetimepad

import "crypto/rand"

// GenerateNewKey generates a new random key using crypto/rand.
func GenerateNewKey(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if isError(err) {
		return nil, err
	}
	return b, nil
}

func isError(err error) bool {
	return err != nil
}
