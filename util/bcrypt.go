package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password.
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password with a cost of 12.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a password with its hashed value and returns true if they match.
func CheckPasswordHash(password, hash string) bool {
	// Compare the password with its hash.
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
