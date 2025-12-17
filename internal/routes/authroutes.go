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
		post.GET("/post/topics", controllers.GetAllTopics)
		post.GET("/post/topic/:topic", controllers.GetPostsByTopic)
		post.GET("/post/id/:id", controllers.GetPostByID)
		post.GET("/post/:author", controllers.GetPostsByAuthor)
		post.DELETE("/deletepost/id/:id", controllers.DeletePostByID)

		// Comments Routes
		post.POST("/comment", controllers.AddComment)
		post.GET("/comments/:postId", controllers.GetCommentsByPostID)
		post.DELETE("/deletecomment/id/:id", controllers.DeleteCommentByID)
	}

}
