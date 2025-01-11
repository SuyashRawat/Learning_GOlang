package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"

)

// HashPassword hashes a plaintext password.
func HashPassword(password string) (string, error) {
	// Generate a hashed password with bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error while hashing password: ", err)
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a hashed password with a plaintext password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
