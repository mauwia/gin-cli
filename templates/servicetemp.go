package templates

// ServerTemplate returns the basic Gin server code as a string.
// It replaces the placeholder with the provided module name.
func ServiceTemplate(moduleName string) string {
	return `package services`
}
