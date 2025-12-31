package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = func() []byte {

	if s := os.Getenv("JWT_SECRET"); s != "" {
		return []byte(s)
	}

	return []byte("default_secret_key")
}()

// custom claim to include username and user id
type CustomClaims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateAccessToken generates a JWT access token for a given username
func GenerateAccessToken(username string, userID uint) (string, error) {
	claims := CustomClaims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseAccessToken parses a JWT access token and returns the username and userID
func ParseAccessToken(tokenString string) (string, uint, error) {
	return parseJWT(tokenString)
}

func GetUsernameFromAccessToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	username, _, err := ParseAccessToken(tokenString)
	return username, err
}

func GetUserIDFromAccessToken(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	_, userID, err := ParseAccessToken(tokenString)
	return userID, err
}

// Generates a refresh token valid for 7 days
func GenerateRefreshToken(username string, userID uint) (string, error) {
	claims := CustomClaims{
		Username: username,
		UserID:   userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseRefreshToken(tokenString string) (string, uint, error) {
	return parseJWT(tokenString)
}

// parses token and returns username and userID
func parseJWT(tokenString string) (string, uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return "", 0, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Username, claims.UserID, nil
	}
	return "", 0, errors.New("invalid token claims")
}
