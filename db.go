package main

import (
	"log"
	"os"
	"split-it/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable(&models.Friend{})
	// db.Migrator().DropTable(&models.Expense{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Friend{})
	db.AutoMigrate(&models.Expense{})

	return db
}
