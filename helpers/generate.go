package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mauwia/gin-cli/templates"
	crudTemplate "github.com/mauwia/gin-cli/templates/crud"
)

func HandleGenerate(arg []string) {
	switch arg[2] {
	case "service":
		CreateService(os.Args[3])
	case "crud":
		CreateCrud(os.Args[3])
	}

}
func CreateCrud(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	// get current directory name and use it as project name
	_, projectName := filepath.Split(cwd)

	CreateFile(filepath.Join(cwd, "internal", "services", fmt.Sprintf("%s.go", name)))
	WriteCrudFile(CapitalizeFirstLetter(&name), &projectName, filepath.Join(cwd, "internal", "services", fmt.Sprintf("%s.go", name)), crudTemplate.CrudServiceTemplate)
	CreateFile(filepath.Join(cwd, "internal", "repositories", fmt.Sprintf("%s.go", name)))
	WriteCrudFile(CapitalizeFirstLetter(&name), &projectName, filepath.Join(cwd, "internal", "repositories", fmt.Sprintf("%s.go", name)), crudTemplate.CrudRepositoryTemplate)
	CreateFile(filepath.Join(cwd, "internal", "controllers", fmt.Sprintf("%s.go", name)))
	WriteCrudFile(CapitalizeFirstLetter(&name), &projectName, filepath.Join(cwd, "internal", "controllers", fmt.Sprintf("%s.go", name)), crudTemplate.CrudControllerTemplate)
	CreateFile(filepath.Join(cwd, "internal", "models", fmt.Sprintf("%s.go", name)))
	WriteFile(CapitalizeFirstLetter(&name), filepath.Join(cwd, "internal", "models", fmt.Sprintf("%s.go", name)), crudTemplate.CrudModelTemplate)
	content, err := os.ReadFile(filepath.Join(cwd, "internal", "server", "router.go"))
	if err != nil {

		fmt.Printf("failed to read router.go: %s", err)
	}
	existingContent := string(content)
	updatedContent := UpdateRouters(existingContent, projectName, name)
	err = os.WriteFile(filepath.Join(cwd, "internal", "server", "router.go"), []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("failed to write updated router.go: %s", err)
	}
	content, err = os.ReadFile(filepath.Join(cwd, "internal", "server", "server.go"))
	if err != nil {

		fmt.Printf("failed to read server.go: %s", err)
	}
	existingContent = string(content)
	updatedContent = addCrudInitializer(existingContent, name)
	err = os.WriteFile(filepath.Join(cwd, "internal", "server", "server.go"), []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("failed to write updated server.go: %s", err)
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
func CapitalizeFirstLetter(s *string) *string {
	if len(*s) == 0 {
		return s
	}
	str := *s
	result := strings.ToUpper(str[:1]) + str[1:]
	return &result
}
func UpdateRouters(content string, name string, moduleName string) string {
	re := regexp.MustCompile(`(?s)import\s*\(\s*(.*?)\s*\)`)
	matches := re.FindStringSubmatch(content)

	if len(matches) > 1 {
		fullImportBlock := matches[0]
		importsContent := matches[1]

		importsToAdd := []string{
			fmt.Sprintf("\"%s/internal/controllers\"", name),
			fmt.Sprintf("\"%s/internal/services\"", name),
			fmt.Sprintf("\"%s/internal/repositories\"", name),
		}

		var updatedImports string = importsContent

		for _, imp := range importsToAdd {
			if !strings.Contains(importsContent, imp) {
				updatedImports += "\n\t" + imp
			}
		}

		// If we added something new, update the block
		if updatedImports != importsContent {
			newImportBlock := fmt.Sprintf("import (\n\t%s\n)", updatedImports)
			content = strings.Replace(content, fullImportBlock, newImportBlock, 1)
		}
	}
	content = strings.TrimSpace(content) + "\n\n" + crudTemplate.CrudRouterTemplate(*CapitalizeFirstLetter(&moduleName), moduleName)
	return content
}
func addCrudInitializer(content string, name string) string {
	cfgRe := regexp.MustCompile(`cfg := config\.LoadConfig\(\)`)
	cfgSection := cfgRe.FindStringIndex(content)
	insertPos := cfgSection[1]
	routerInit := fmt.Sprintf("\n\tSetup%sRoutes(r,cfg)\n", *CapitalizeFirstLetter(&name))
	content = content[:insertPos] + routerInit + content[insertPos:]
	return content
}
