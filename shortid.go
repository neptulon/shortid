package shortid

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math"
)

// ID generates a base64 encoded URL safe unique ID using crypto/rand.
// This is a cryptographically secure pseudorandom number presented
// as an unpadded base64 encoded string.
//
// 'bits' should be a multiple of 8.
// Lenght of generated string will be ceiling(bits * 4 / 24)
func ID(bits int) (string, error) {
	if math.Mod(float64(bits), 8) != 0 {
		return "", errors.New("Given bits should be a multiple of 8")
	}

	b := make([]byte, bits/8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

// UUID generates a base64 encoded URL safe unique ID using crypto/rand.
// As per UUID v4 specs, this is a 128 bits cryptographically secure
// pseudorandom number, presented as an unpadded base64 encoded string.
func UUID() (string, error) {
	return ID(128)
}
