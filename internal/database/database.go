package database

import (
	"fmt"
	"log"

	"github.com/fitnis/patient-service/configs"
	"github.com/fitnis/patient-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// ConnectDB connects to the database and performs migrations
func ConnectDB(config *configs.Config) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Connected to the database successfully")

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Patient{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	DB = db
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
