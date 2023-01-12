package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// GenerateToken generates a JWT token for the given user ID.
func GenerateToken(userID string, secret string) (string, error) {
	// Set the token claims.
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   userID,
	}

	// Create the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and return it.
	return token.SignedString([]byte(secret))
}

// HashPassword hashes a password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
