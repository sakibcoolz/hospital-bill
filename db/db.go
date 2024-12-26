package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnect() *gorm.DB {
	// Database connection string (replace with your actual credentials)
	dsn := "host=0.0.0.0 user=postgres dbname=postgres port=5432 sslmode=disable"

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
