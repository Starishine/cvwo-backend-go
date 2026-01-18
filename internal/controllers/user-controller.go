package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Starishine/cvwo-backend-go/internal/database"
	"github.com/Starishine/cvwo-backend-go/internal/models"
	"github.com/Starishine/cvwo-backend-go/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists. Please choose a different username."})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login and token generation
func LoginUser(c *gin.Context) {
	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	// Check if user exists with given username and password
	result := database.DB.Where("username = ? AND password = ?", body.Username, body.Password).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT access token
	accessToken, err := utils.GenerateAccessToken(user.Username, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken(user.Username, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Use raw header to set cookie
	c.Writer.Header().Set("Set-Cookie",
		fmt.Sprintf("refresh_token=%s; Path=/; Max-Age=%d; HttpOnly; Secure; SameSite=None",
			refreshToken, 7*24*60*60))

	// return access token in JSON response
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "access_token": accessToken})

}

// RefreshToken handles token refreshing
func RefreshToken(c *gin.Context) {
	log.Println("=== REFRESH TOKEN ENDPOINT CALLED ===")

	// Log ALL cookies
	cookies := c.Request.Cookies()
	log.Printf("All cookies received: %d", len(cookies))
	for _, c := range cookies {
		log.Printf("Cookie: %s = %s", c.Name, c.Value)
	}

	// get refresh token from cookie
	cookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		log.Println("ERROR: No refresh token cookie found")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No refresh token provided"})
		return
	}

	log.Printf("Found refresh token: %s", cookie.Value)

	// validate refresh token and get username and userId
	username, _, err := utils.ParseRefreshToken(cookie.Value)
	_, userId, err := utils.ParseRefreshToken(cookie.Value)

	if err != nil {
		log.Println("ERROR: Invalid refresh token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	log.Printf("Valid token for user: %s", username)
	log.Printf("UserId: %d", userId)

	// generate new access token
	newAccessToken, err := utils.GenerateAccessToken(username, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new token"})
		return
	}

	// gererate new refresh token
	refreshToken, err := utils.GenerateRefreshToken(username, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Use raw header to set cookie
	c.Writer.Header().Set("Set-Cookie",
		fmt.Sprintf("refresh_token=%s; Path=/; Max-Age=%d; HttpOnly; Secure; SameSite=None",
			refreshToken, 7*24*60*60))

	// return access token in JSON response
	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})

}
