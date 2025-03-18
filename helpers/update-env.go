package helpers

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindUnusedVars(existingVars, envVars []string) []string {
	envSet := make(map[string]struct{})
	for _, envVar := range envVars {
		envSet[envVar] = struct{}{}
	}

	var unused []string
	for _, field := range existingVars {
		if _, found := envSet[field]; !found {
			unused = append(unused, field)
		}
	}
	return unused
}

func RemoveUnusedVarsFromConfig(content string, unusedVars []string) string {
	// Remove from struct
	for _, field := range unusedVars {
		fieldPattern := fmt.Sprintf(`\s+%s\s+string\n?`, field)
		re := regexp.MustCompile(fieldPattern)
		content = re.ReplaceAllString(content, "\n")
	}

	// Remove from LoadConfig
	for _, field := range unusedVars {
		loadConfigPattern := fmt.Sprintf(`\t\t%s:\s+getEnv\("[^"]+",\s*"[^"]*"\),?\n?`, field)
		re := regexp.MustCompile(loadConfigPattern)
		content = re.ReplaceAllString(content, "")
	}

	return content
}

func ReadEnvFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var vars []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			vars = append(vars, strings.TrimSpace(parts[0]))
		}
	}
	return vars, scanner.Err()
}
func ExtractExistingConfigVars(content string) []string {
	re := regexp.MustCompile(`type AppConfig struct \{([^}]+)\}`)
	matches := re.FindStringSubmatch(content)

	var existingVars []string
	if len(matches) > 1 {
		fieldsBlock := matches[1]
		fieldRe := regexp.MustCompile(`\s+([A-Za-z0-9_]+)\s+string`)
		fieldMatches := fieldRe.FindAllStringSubmatch(fieldsBlock, -1)

		for _, match := range fieldMatches {
			existingVars = append(existingVars, match[1])
		}
	}
	return existingVars

}
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false

}
func UpdateEnvGo() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}

	content, err := os.ReadFile(filepath.Join(cwd, "config", "config.go"))
	if err != nil {

		return fmt.Errorf("failed to read config.go: %w", err)
	}
	existingContent := string(content)

	envVars, err := ReadEnvFile(filepath.Join(cwd, ".env"))

	if err != nil {
		return fmt.Errorf("failed to read .env: %w", err)
	}

	existingVars := ExtractExistingConfigVars(existingContent)
	var newVars []string
	for _, envVar := range envVars {
		if !Contains(existingVars, envVar) {
			newVars = append(newVars, envVar)
		}
	}
	// if len(newVars) == 0 {
	// 	return nil
	// }
	unusedVars := FindUnusedVars(existingVars, envVars)
	updatedContent := AddNewVarsToConfig(existingContent, newVars)
	updatedContent = RemoveUnusedVarsFromConfig(updatedContent, unusedVars)
	err = os.WriteFile(filepath.Join(cwd, "config", "config.go"), []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated config.go: %w", err)
	}

	fmt.Println("config.go updated successfully!")
	return nil
}
func AddNewVarsToConfig(content string, newVars []string) string {
	structRe := regexp.MustCompile(`type AppConfig struct {`)
	structSection := structRe.FindStringIndex(content)
	insertPos := structSection[1]
	for _, v := range newVars {
		content = content[:insertPos] + "\n\t" + v + " string" + content[insertPos:]
	}
	loadRe := regexp.MustCompile(`return &AppConfig\{`)
	loadSection := loadRe.FindStringIndex(content)
	loadInsertPos := loadSection[1]
	for _, v := range newVars {
		content = content[:loadInsertPos] + "\n\t\t" + v + ": getEnv(\"" + v + "\", \"\")," + content[loadInsertPos:]
	}

	return content

}
