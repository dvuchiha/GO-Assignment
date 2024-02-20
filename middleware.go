package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Log the incoming request & response
func logRequest(c *gin.Context) {
	fmt.Printf("[%s] %s - %s %s\n", time.Now().Format("2006-01-02 15:04:05"), c.ClientIP(), c.Request.Method, c.Request.URL.Path)
	// fmt.Printf("Request Headers: %v\n", c.Request.Header)

	c.Next()

	fmt.Printf("[%s] %d %s\n", time.Now().Format("2006-01-02 15:04:05"), c.Writer.Status(), http.StatusText(c.Writer.Status()))
	// fmt.Printf("Response Headers: %v\n", c.Writer.Header())

}
