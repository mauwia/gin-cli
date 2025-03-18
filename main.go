package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mauwia/gin-cli/helpers"
	"github.com/mauwia/gin-cli/templates"
)

// Define a template for a basic Gin server.

func handleCase4(args []string) {
	switch args[1] {
	case "generate":
		println("generate")
		helpers.HandleGenerate(args)
	}
}
func handleCase3(args []string) {
	switch args[1] {
	case "new":
		initProject(&args[2])
	case "update":
		if args[2] != "env" {
			fmt.Fprintf(os.Stderr, "Usage: %s <project_name> [new|generate]\n", os.Args[0])
			os.Exit(1)
		}
		err := helpers.UpdateEnvGo()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update config.go: %v\n", err)
			os.Exit(1)
		}
	case "setup":
		helpers.HandleSetup(args)
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s <project_name> [new|generate]\n", os.Args[0])
		os.Exit(1)
	}

}
func initProject(outputDir *string) {
	// Define flags to customize the output directory and file name.

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}

	// Use the executable directory as the base for the output directory.
	helpers.CreateFolder(filepath.Join(cwd, *outputDir))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "cmd"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "config"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "migrations"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "server"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "handlers"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "models"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "services"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "repositories"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "middlewares"))
	helpers.CreateFolder(filepath.Join(cwd, *outputDir, "internal", "utils"))

	// Create the output directory if it doesn't exist.

	// Build the full path for the output file.
	helpers.CreateFile(filepath.Join(*outputDir, ".env"))

	helpers.CreateFile(filepath.Join(*outputDir, "main.go"))
	helpers.CreateFile(filepath.Join(*outputDir, "config", "config.go"))
	helpers.CreateFile(filepath.Join(*outputDir, "internal", "server", "server.go"))
	helpers.CreateFile(filepath.Join(*outputDir, "internal", "server", "router.go"))
	// Create the output file.

	// Parse and execute the template, writing the contents to the file.
	helpers.WriteFile(outputDir, filepath.Join(*outputDir, ".env"), templates.ENVTemplate)

	helpers.WriteFile(outputDir, filepath.Join(*outputDir, "main.go"), templates.MainTemplate)
	helpers.WriteFile(outputDir, filepath.Join(*outputDir, "internal", "server", "server.go"), templates.ServerTemplate)
	helpers.WriteFile(outputDir, filepath.Join(*outputDir, "internal", "server", "router.go"), templates.RouterTemplate)
	helpers.WriteFile(outputDir, filepath.Join(*outputDir, "config", "config.go"), templates.ConfigTemplate)

	helpers.InitGin(outputDir)

	fmt.Printf("Go module initialized in %s with module name \n", *outputDir)
	fmt.Printf("Basic Gin server code generated in")
}
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <project_name>\n", os.Args[0])
		os.Exit(1)
	}
	switch len(os.Args) {
	case 3:
		handleCase3(os.Args)
	case 4:
		handleCase4(os.Args)
	}

}
