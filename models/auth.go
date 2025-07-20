package models

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

// AuthCredentials represents the data needed for authentication
// For registration and login, users must provide email and password
// During registration, users can select a role: either "manager" or "attendee"
// If no role is provided during registration, 'attendee' will be set as the default role
type AuthCredentials struct {
	Email    string   `json:"email" validate:"required" example:"user@example.com"`
	Password string   `json:"password" validate:"required" example:"securepassword"`
	Role     UserRole `json:"role,omitempty" example:"attendee" enums:"manager,attendee"` // Optional: must be either "manager" or "attendee"
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *AuthCredentials) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type AuthService interface {
	Login(ctx context.Context, loginData *AuthCredentials) (string, *User, error)
	Register(ctx context.Context, registerData *AuthCredentials) (string, *User, error)
}

// Check if a password matches a hash
func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Checks if an email is valid
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
