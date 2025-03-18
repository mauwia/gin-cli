package templates

// ServerTemplate returns the basic Gin server code as a string.
// It replaces the placeholder with the provided module name.
func RouterTemplate(moduleName string) string {
	return `package server

import (
	"github.com/gin-gonic/gin"
)

// setupRoutes initializes the routes
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", pingHandler)
	}
}

// pingHandler handles the /ping route
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


`
}
