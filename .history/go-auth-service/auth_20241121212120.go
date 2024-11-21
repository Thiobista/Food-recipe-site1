// auth.go
package main

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your_secret_key_here")

// GenerateJWT generates a new JWT token for a user.
func GenerateJWT(userID string) (string, error) {
	// Define claims
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Token expiration
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(secretKey)
}

// ValidateJWT validates the JWT token.
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secretKey, nil
	})
	return token, err
}
