package main

import (
	"fmt"
	"log"
	"os"
	"split-it/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Friend{})
	db.Migrator().DropTable(&models.Expense{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Friend{})
	db.AutoMigrate(&models.Expense{})

	return db
}
