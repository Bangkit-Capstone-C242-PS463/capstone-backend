package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"capstone-backend/internal/constants"
)

var jwtSecretKey = []byte(os.Getenv("ACCESS_SECRET"))

// Generates a JWT token for a given user ID
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		constants.CONTEXT_USERID_KEY: userID,
		"exp":                        time.Now().Add(time.Hour * 4).Unix(), // Expiration time (4 hour)
		"iat":                        time.Now().Unix(),                    // Issued at
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// Verifies the JWT token and returns the claims if valid
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
