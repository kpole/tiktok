package utils

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

func MD5(str string) (string, error) {
	h := md5.New()
	if _, err := io.WriteString(h, str); err != nil {
		return "", err
	}
	data := fmt.Sprintf("%x", h.Sum(nil))
	return data, nil
}

func crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	cost := 5

	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

func verifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
