package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func InitGin(outputDir *string) {
	cmd := exec.Command("go", "mod", "init", *outputDir)
	cmd.Dir = *outputDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run 'go mod init': %v\nOutput: %s\n", err, output)
		os.Exit(1)
	}
	installCmd := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin", "github.com/joho/godotenv")
	fmt.Printf("Installing Gin package...\n")
	installCmd.Dir = *outputDir
	_, err = installCmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run 'go get -u': %v\n", err)
		os.Exit(1)
	}

}
func CreateFile(outputPath string) {
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
}
func WriteFile(outputDir *string, outputPath string, templ func(string) string) {
	f, _ := os.Create(outputPath)
	tmpl, err := template.New("server").Parse(templ(*outputDir))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse template: %v\n", err)
		os.Exit(1)
	}

	if err := tmpl.Execute(f, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to file: %v\n", err)
		os.Exit(1)
	}
}
func CreateFolder(finalOutputDir string) {
	err := os.MkdirAll(finalOutputDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create directory: %v\n", err)
		os.Exit(1)
	}
}
