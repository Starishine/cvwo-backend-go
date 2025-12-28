package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUser(c *gin.Context) {
	log.Println("=== LOGOUT ENDPOINT CALLED ===")

	cookies := c.Request.Cookies()
	log.Printf("Cookies received: %d", len(cookies))
	for _, cookie := range cookies {
		log.Printf("Cookie: %s = %s", cookie.Name, cookie.Value)
	}

	// Using raw header to clear cookie
	c.Writer.Header().Add("Set-Cookie", "refresh_token=; Path=/; Max-Age=0; HttpOnly; SameSite=Lax")
	c.Writer.Header().Add("Set-Cookie", "refresh_token=; Path=/; Domain=localhost; Max-Age=0; HttpOnly; SameSite=Lax")

	log.Println("Sent Set-Cookie headers to clear cookie")

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
