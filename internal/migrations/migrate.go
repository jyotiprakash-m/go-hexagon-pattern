package migrations

import (
	"log"

	"gin-hex/internal/core/models"

	"gorm.io/gorm"
)

// RunMigrations applies the database migrations
func RunMigrations(db *gorm.DB) {
	log.Println("Running Migrations...")

	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("✅ Database Migration Completed!")
}
