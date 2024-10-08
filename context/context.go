package context

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route that triggers an HTTP request with context
	router.GET("/proxy", func(c *gin.Context) {
		// Set a timeout for the context (e.g., 3 seconds)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel() // Ensure that the context is canceled to avoid resource leaks

		// Create a new HTTP request with the context
		req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/5", nil)
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Make the HTTP request using the default client
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// Handle context timeout or request error
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("Request timed out")
				c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timed out"})
				return
			}
			log.Printf("Request failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Request failed"})
			return
		}
		defer resp.Body.Close()

		// Handle the HTTP response
		if resp.StatusCode != http.StatusOK {
			c.JSON(resp.StatusCode, gin.H{"error": "Upstream request failed"})
			return
		}

		// Send a success response
		c.JSON(http.StatusOK, gin.H{"message": "Request successful", "status": resp.Status})
	})

	// Start the Gin server
	router.Run(":8080")
}
