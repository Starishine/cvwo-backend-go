package router

import (
	"github.com/Starishine/cvwo-backend-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	routes.AuthRoutes(r)
	return r
}
