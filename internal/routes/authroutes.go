package routes

import (
	"github.com/Starishine/cvwo-backend-go/internal/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}

	post := r.Group("/")
	{
		post.POST("/post", controllers.AddPost)
	}

}
