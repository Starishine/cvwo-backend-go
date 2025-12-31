package controllers

import (
	"fmt"
	"strconv"

	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/Starishine/cvwo-backend-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikePost(c *gin.Context) {

	// Get the userId
	userId, err := utils.GetUserIDFromAccessToken(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	fmt.Println("User ID from token:", userId)

	// parse postId
	postId, err := strconv.ParseUint(c.Param("postId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}
	postIdUint := uint(postId)

	// check if post exists
	var post models.Post
	if err := database.DB.First(&post, postIdUint).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	// initially set to false to trackk if like or unlike
	wasLiked := false

	// using transaction to ensure data consistency
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var existingLike models.Like
		// find like for this post by this user
		result := tx.Where("post_id = ? AND user_id = ?", postIdUint, userId).First(&existingLike)

		if result.Error == nil {
			// Like already exists and is active, so unlike the post
			wasLiked = false
			if err := tx.Unscoped().Delete(&existingLike).Error; err != nil {
				return err
			}

			// Decrement like count on post
			if err := tx.Model(&models.Post{}).Where("id = ?", postIdUint).UpdateColumn("likes", gorm.Expr("GREATEST(likes - 1, 0)")).Error; err != nil {
				return err
			}

			return nil
		}

		// No like record at all -> create a new like
		wasLiked = true

		newLike := models.Like{
			UserID: userId,
			PostID: postIdUint,
		}

		if err := tx.Create(&newLike).Error; err != nil {
			return err
		}

		// Increment like count on post
		if err := tx.Model(&models.Post{}).Where("id = ?", postIdUint).Update("likes", gorm.Expr("likes + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to like/unlike post"})
		return
	}

	// Get updated post with like count
	var updatedPost models.Post
	database.DB.First(&updatedPost, postIdUint)

	if wasLiked {
		c.JSON(200, gin.H{
			"message": "Post liked",
			"liked":   true,
			"likes":   updatedPost.Likes,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Post unliked",
			"liked":   false,
			"likes":   updatedPost.Likes,
		})
	}

}
