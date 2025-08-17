package config

import (
	"fmt"
	"github.com/filtkac/consult-scheduler/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func DatabaseConnection() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable timezone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.ConsultType{},
		&model.Department{},
		&model.DayTemplate{},
		&model.DayTemplateConsult{},
		&model.Patient{},
		&model.Consult{},
	)

	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	log.Println("Database migrations complete.")
}
