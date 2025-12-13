package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Starishine/cvwo-backend-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	fmt.Println("HOST", os.Getenv("DB_HOST"))
	fmt.Println("USER:", os.Getenv("DB_USER"))
	fmt.Println("PASS:", os.Getenv("DB_PASSWORD"))
	fmt.Println("NAME:", os.Getenv("DB_NAME"))
	fmt.Println("PORT:", os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = db

	// Auto-migrate the database schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Comment{})

	log.Println("Database is connected")
}
