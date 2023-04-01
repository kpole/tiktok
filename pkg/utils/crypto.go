package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	cost := 5

	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
