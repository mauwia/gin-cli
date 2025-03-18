package templates

import "fmt"

// MainTemplate returns the basic Gin server code as a string.
// It replaces the placeholder with the provided module name.
func MainTemplate(moduleName string) string {
	return fmt.Sprintf(`package main

import (
    "log"
    "%s/internal/server"
)

func main() {

    s := server.NewServer()
    if err := s.Run(); err != nil {
        log.Fatal(err)
    }
}
`, moduleName)
}
