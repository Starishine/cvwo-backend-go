package controllers

import (
	"net/http"

	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Table("posts").Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post added successfully"})
}

func GetPostsByAuthor(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Table("posts").Find(&posts, "author = ?", c.Param("author"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetAllTopics(c *gin.Context) {
	var topics []string
	result := database.DB.Table("posts").Distinct().Pluck("topic", &topics)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func GetPostsByTopic(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Table("posts").Find(&posts, "topic = ?", c.Param("topic"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
