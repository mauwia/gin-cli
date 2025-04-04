package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	configTemplates "github.com/mauwia/gin-cli/templates/config"
)

func HandleSetup(arg []string) {
	switch arg[2] {
	case "postgres":
		setupPostgres()
	case "mongodb":
		setupMongoDB()
	}

}
func setupPostgres() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Setting up Postgres...")
	cmd := exec.Command("go", "get", "-u", "gorm.io/gorm", "gorm.io/driver/postgres")
	cmd.Dir = cwd
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run 'go get': %v\nOutput: %s\n", err, output)
		os.Exit(1)
	}
	fmt.Println("Postgres dependencies installed successfully!")
	content, err := os.ReadFile(filepath.Join(cwd, ".env"))
	if err != nil {
		return fmt.Errorf("failed to read config.go: %w", err)
	}
	envVars := string(content)
	envVars += "\nDB_USER=postgres\n"
	envVars += "DB_PASSWORD=password\n"
	envVars += "DB_NAME=postgres\n"
	envVars += "DB_PORT=5432\n"
	envVars += "DB_HOST=localhost\n"
	err = os.WriteFile(filepath.Join(cwd, ".env"), []byte(envVars), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated config.go: %w", err)
	}
	UpdateEnvGo()
	CreateFile(filepath.Join(cwd, "config", "postgresdbconfig.go"))
	WriteFile(&cwd, filepath.Join(cwd, "config", "postgresdbconfig.go"), configTemplates.PostgresDBConfigTemplate)
	content, err = os.ReadFile(filepath.Join(cwd, "internal", "server", "server.go"))
	if err != nil {

		return fmt.Errorf("failed to read config.go: %w", err)
	}
	existingContent := string(content)
	updatedContent := addDBInitializer(existingContent, "SetupPostgresDB")
	err = os.WriteFile(filepath.Join(cwd, "internal", "server", "server.go"), []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated config.go: %w", err)
	}

	fmt.Println("Postgres setup successfully!")
	return nil
}
func setupMongoDB() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Setting up Mongodb...")
	cmd := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin", "go.mongodb.org/mongo-driver/mongo", "github.com/joho/godotenv")
	cmd.Dir = cwd
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run 'go get': %v\nOutput: %s\n", err, output)
		os.Exit(1)
	}
	fmt.Println("Mongodb dependencies installed successfully!")
	content, err := os.ReadFile(filepath.Join(cwd, ".env"))
	if err != nil {
		return fmt.Errorf("failed to read config.go: %w", err)
	}
	envVars := string(content)
	envVars += "\nMONGODB_URI=\n"
	envVars += "MONGODB_DATABASE=\n"
	err = os.WriteFile(filepath.Join(cwd, ".env"), []byte(envVars), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated config.go: %w", err)
	}
	CreateFile(filepath.Join(cwd, "config", "mongodbconfig.go"))
	WriteFile(&cwd, filepath.Join(cwd, "config", "mongodbconfig.go"), configTemplates.MongoDBConfigTemplate)
	UpdateEnvGo()
	content, err = os.ReadFile(filepath.Join(cwd, "internal", "server", "server.go"))
	if err != nil {

		return fmt.Errorf("failed to read config.go: %w", err)
	}
	existingContent := string(content)
	updatedContent := addDBInitializer(existingContent, "SetupMongoDB")
	err = os.WriteFile(filepath.Join(cwd, "internal", "server", "server.go"), []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated config.go: %w", err)
	}
	fmt.Println("MongoDB setup successfully!")
	return nil
}
func addDBInitializer(content string, funcName string) string {
	cfgRe := regexp.MustCompile(`cfg := config\.LoadConfig\(\)`)
	cfgSection := cfgRe.FindStringIndex(content)
	insertPos := cfgSection[1]
	dbInit := fmt.Sprintf("\n\tdb := config.%s(cfg)\n\t println(db)\n", funcName)
	content = content[:insertPos] + dbInit + content[insertPos:]
	return content
}
