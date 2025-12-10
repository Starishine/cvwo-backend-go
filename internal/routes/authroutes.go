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
		post.GET("/post/:author", controllers.GetPostsByAuthor)
		post.GET("/post/topics", controllers.GetAllTopics)
		post.GET("/post/topic/:topic", controllers.GetPostsByTopic)
	}

}
