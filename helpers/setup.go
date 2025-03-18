package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func HandleSetup(arg []string) {
	switch arg[2] {
	case "postgres":
		setupPostgres()
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
	envVars += "\nDB_USER=\n"
	envVars += "DB_PASSWORD=\n"
	envVars += "DB_NAME=\n"
	envVars += "DB_PORT=\n"
	envVars += "DB_HOST=\n"
	err = os.WriteFile(filepath.Join(cwd, ".env"), []byte(envVars), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated config.go: %w", err)
	}
	UpdateEnvGo()
	fmt.Println("Postgres setup successfully!")
	return nil
}
