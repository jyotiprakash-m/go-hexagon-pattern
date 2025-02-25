package db

import (
	"gin-hex/internal/migrations"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB initializes the PostgreSQL database
func ConnectDB() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate the Todo model
    migrations.RunMigrations(db)
    if err != nil {
        log.Fatal("Failed to run migrations:", err)
    }

    log.Println("Database connected & migrated successfully!")
    return db
}
