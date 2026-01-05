package main

import (
	"os"

	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/Starishine/cvwo-backend-go/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// setup DB
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	// initialise router
	r := router.SetupRouter()

	// use dynamic port provided by render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
