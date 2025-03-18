package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mauwia/gin-cli/templates"
)

func HandleGenerate(arg []string) {
	switch arg[2] {
	case "service":
		CreateService(os.Args[3])
	}

}
func CreateService(serviceName string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	CreateFile(filepath.Join(cwd, "internal", "services", fmt.Sprintf("%s.go", serviceName)))
	WriteFile(&cwd, filepath.Join(cwd, "internal", "services", fmt.Sprintf("%s.go", serviceName)), templates.ServiceTemplate)
}
