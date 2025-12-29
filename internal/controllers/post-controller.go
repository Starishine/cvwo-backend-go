package controllers

import (
	"net/http"

	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/Starishine/cvwo-backend-go/internal/utils"
	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Model(&models.Post{}).Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post added successfully"})
}

func GetPostsByAuthor(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Model(&models.Post{}).Find(&posts, "author = ?", c.Param("author"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// Get all the topics from all posts
func GetAllTopics(c *gin.Context) {
	var topics []string
	result := database.DB.Model(&models.Post{}).Distinct().Pluck("topic", &topics)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func GetPostsByTopic(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Model(&models.Post{}).Find(&posts, "topic = ?", c.Param("topic"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	var post models.Post
	result := database.DB.Model(&models.Post{}).First(&post, "id = ?", c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func DeletePostByID(c *gin.Context) {
	username, err := utils.GetUsernameFromAccessToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or missing token"})
		return
	}

	var post models.Post

	result := database.DB.Model(&models.Post{}).First(&post, "id = ?", c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if the current user is the author
	if post.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own posts"})
		return
	}

	delresult := database.DB.Model(&models.Post{}).Delete(&models.Post{}, "id = ?", c.Param("id"))
	if delresult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": delresult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	result := database.DB.Model(&models.Post{}).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
