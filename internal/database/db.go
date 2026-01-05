package database

import (
	"log"
	"os"

	"github.com/Starishine/cvwo-backend-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DATABASE_URL") // pulls url from renders settings

	if dsn == "" {
		// used for your local development
		dsn = "host=localhost user=postgres password=... dbname=cvwo_db port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = db

	// Auto-migrate the database schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Comment{})
	DB.AutoMigrate(&models.Like{})

	log.Println("Database is connected")
}
