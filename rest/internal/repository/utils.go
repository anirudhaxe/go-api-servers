package repository

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

/*
Util function to validate password
*/
func (u *User) ValidatePassword(pw string) bool {
	// casting the string (immutable) to raw byte sequence (mutable) for crypto operations
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}

/*
Util function to generate password hash
*/
func GeneratePasswordHash(pw string) (string, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(encpw), nil

}

/*
Util function to generate JWT
*/
func (u *User) GenerateJwtToken() (string, error) {

	claims := &jwt.MapClaims{
		"username":  u.Username,
		"email":     u.Email,
		"role":      u.Role,
		"is_active": u.IsActive,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(secret))

}

/*
Util function to validate JWT
*/
func ValidateJwtToken(tokenString string) (*jwt.Token, error) {

	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validating the expected algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

}
