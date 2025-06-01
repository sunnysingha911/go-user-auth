package database

import (
	"fmt"
	"log"

	"github.com/sunnysingha911/user-service/config"
	"github.com/sunnysingha911/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	// Enable uuid-ossp extension
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		log.Fatal("Failed to enable uuid-ossp extension: ", err)
	}

	// Run migrations
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run DB migration: ", err)
	}

	DB = db

	fmt.Println("Connected to PostgreSQL")

	return nil
}
