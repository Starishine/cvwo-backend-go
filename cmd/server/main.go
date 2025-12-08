package main

import (
	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/Starishine/cvwo-backend-go/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{})

	r := router.SetupRouter()
	r.Run(":8080")
}
