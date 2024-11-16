package utils

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
)

// Creates a random salt of the specified length as a []byte
func GenerateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Hashes a password with a provided salt (as []byte) using bcrypt
func HashPasswordWithSalt(password string, salt []byte) (string, error) {
	passwordWithSalt := append([]byte(password), salt...)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Generates a hashed password along with its salt
func CreateHashedPassword(password string, saltLength int) (string, []byte, error) {
	salt, err := GenerateSalt(saltLength)
	if err != nil {
		return "", nil, err
	}

	hashedPassword, err := HashPasswordWithSalt(password, salt)
	if err != nil {
		return "", nil, err
	}

	return hashedPassword, salt, nil
}

// Verifies a plain-text password against a hashed password and salt
func VerifyPasswordWithSalt(hashedPassword string, plainPassword string, salt []byte) bool {
	passwordWithSalt := append([]byte(plainPassword), salt...)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), passwordWithSalt)
	return err == nil
}
