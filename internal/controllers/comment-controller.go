package controllers

import (
	"fmt"
	"net/http"

	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/Starishine/cvwo-backend-go/internal/utils"
	"github.com/gin-gonic/gin"
)

// adds comment to a post and saves to database
func AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Model(&models.Comment{}).Create(&comment)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment added successfully"})
}

func GetCommentsByPostID(c *gin.Context) {
	var comments []models.Comment
	result := database.DB.Model(&models.Comment{}).Where("post_id = ?", c.Param("postId")).
		Where("parent_id IS NULL OR parent_id = 0").
		Order("created_at DESC").Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Println(comments)
	c.JSON(http.StatusOK, comments)
}

func DeleteCommentByID(c *gin.Context) {
	username, err := utils.GetUsernameFromAccessToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or missing token"})
		return
	}

	var comment models.Comment
	result := database.DB.Model(&models.Comment{}).First(&comment, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if comment.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own comments"})
		return
	}

	deleteRes := database.DB.Delete(&comment)
	if deleteRes.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": deleteRes.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

// retrieve replies for a specific comment
func GetRepliesByCommentID(c *gin.Context) {
	var replies []models.Comment
	result := database.DB.Model(&models.Comment{}).Where("parent_id = ?", c.Param("parentId")).
		Order("created_at ASC").Find(&replies)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"replies": replies})
}
