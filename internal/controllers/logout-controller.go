package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUser(c *gin.Context) {
	log.Println("=== LOGOUT ENDPOINT CALLED ===")

	// Log request headers to see what browser is sending
	log.Println("=== ALL REQUEST HEADERS ===")
	for name, values := range c.Request.Header {
		for _, value := range values {
			log.Printf("Header: %s = %s", name, value)
		}
	}

	cookies := c.Request.Cookies()
	log.Printf("Cookies received: %d", len(cookies))
	for i, cookie := range cookies {
		log.Printf("Cookie %d: Name=%s, Value=%s, Domain=%s, Path=%s, Secure=%v, HttpOnly=%v",
			i, cookie.Name, cookie.Value, cookie.Domain, cookie.Path, cookie.Secure, cookie.HttpOnly)
	}

	// Try to clear with EVERY possible combination
	variations := []struct{ domain, path string }{
		{"localhost", "/"},
		{"", "/"},
		{"localhost", "/auth"},
		{"", "/auth"},
	}

	for _, v := range variations {
		c.SetCookie(
			"refresh_token", // name
			"",              // value
			-1,              // maxAge (forces deletion)
			"/",             // path
			"localhost",     // domain
			false,           // secure (set to true if using HTTPS)
			true,            // httpOnly
		)
		log.Printf("Attempted clear: domain='%s', path='%s'", v.domain, v.path)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})

}
