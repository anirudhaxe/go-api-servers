package model

import "github.com/anirudhaxe/go-api-servers/rest/internal/repository"

type UserRole = repository.UserRole

// User model
type User struct {
	ID                string   `json:"id"`
	Username          string   `json:"username"`
	Email             string   `json:"email"`
	EncryptedPassword string   `json:"-"` // Don't include in JSON
	Role              UserRole `json:"role"`
	IsActive          bool     `json:"is_active"`
	CreatedAt         string   `json:"created_at"`
	LastLoginAt       *string  `json:"last_login_at,omitempty"`
}

// Register user request
type RegisterUserRequest struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

// Login request
type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response for both login/register requests
type UserResponse struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Role     UserRole `json:"role"`
	Token    string   `json:"token"`
}
