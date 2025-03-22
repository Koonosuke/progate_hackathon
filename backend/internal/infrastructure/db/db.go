package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v\n", err)
		}
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\nConnection String: %s\n", err, url)
	}

	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("Failed to get database connection:", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalln("Failed to close database connection:", err)
	}
}
