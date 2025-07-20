package db

import (
	"github.com/aurareaper/event-management-app/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	// Check if users table exists before trying to fix the created_at column
	if db.Migrator().HasTable(&models.User{}) {
		// Fix empty string values in created_at column to avoid timestamp conversion errors
		// First set empty strings to NULL
		db.Exec("UPDATE users SET created_at = NULL WHERE created_at = ''")
		db.Exec("UPDATE users SET updated_at = NULL WHERE updated_at = ''")
		
		// Then handle the timestamp conversion separately
		db.Exec("ALTER TABLE users ALTER COLUMN created_at TYPE timestamptz USING created_at::timestamptz")
		db.Exec("ALTER TABLE users ALTER COLUMN updated_at TYPE timestamptz USING updated_at::timestamptz")
	}
	
	// Perform the auto-migration
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}
