package helpers

import (
	"fmt"
	"os"
)

// ShowHelp displays the CLI documentation and usage information
func ShowHelp() {
	fmt.Print(`
Gin CLI - A tool for scaffolding Go web applications using the Gin framework

USAGE:
  gin-cli COMMAND [ARGUMENTS]

COMMANDS:
  new <project-name>             Create a new Gin project with clean architecture
  generate crud <entity-name>    Generate CRUD operations for an entity
  generate service <name>        Generate a new service
  setup postgres                 Set up PostgreSQL database configuration
  setup mongodb                  Set up MongoDB database configuration
  update env                     Update environment configuration

EXAMPLES:
  gin-cli new my-api             Create a new project named "my-api"
  gin-cli generate crud user     Generate CRUD operations for "user" entity
  gin-cli generate service auth  Generate a new "auth" service
  gin-cli setup postgres         Configure PostgreSQL for your project
  gin-cli update env             Update config.go based on your .env file

For more information, visit: https://github.com/mauwia/gin-cli
`)
	os.Exit(1)
}
