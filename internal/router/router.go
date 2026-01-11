package router

import (
	"os"
	"time"

	"github.com/Starishine/cvwo-backend-go/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // Fallback for local dev
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:        []string{frontendURL},
		AllowMethods:        []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:        []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:       []string{"Content-Length", "Set-Cookie"},
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
		MaxAge:              12 * time.Hour,
	}))
	routes.AuthRoutes(r)
	return r
}
