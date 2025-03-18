package templates

import "fmt"

// ServerTemplate returns the basic Gin server code as a string.
// It replaces the placeholder with the provided module name.
func ServerTemplate(moduleName string) string {
	return fmt.Sprintf(`package server

import (
	"github.com/gin-gonic/gin"
	"%s/config"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	SetupRoutes(r)

	return &Server{router: r}
}

func (s *Server) Run() error {
    cfg := config.LoadConfig()
	return s.router.Run(":" + cfg.PORT)
}


`, moduleName)
}
