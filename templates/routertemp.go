package templates

import "fmt"

func RouterTemplate(moduleName string) string {
	return fmt.Sprintf(`package server

import (
	"github.com/gin-gonic/gin"
	"%s/config"
)

// setupRoutes initializes the routes
func SetupRoutes(r *gin.Engine,cfg *config.AppConfig) {
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


`, moduleName)
}
