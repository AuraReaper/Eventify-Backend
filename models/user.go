package models

import (
	"time"
)

// UserRole defines the type of user account
type UserRole string

// Available user roles in the system
const (
	Manager  UserRole = "manager"  // Manager role with event management privileges
	Attendee UserRole = "attendee" // Attendee role for booking tickets to events
)

// RoleOptions returns a list of valid user roles for documentation and validation
func RoleOptions() []string {
	return []string{string(Manager), string(Attendee)}
}

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"text;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:attendee"`
	Password  string    `json:"-"` // Do not compute the password in json
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
