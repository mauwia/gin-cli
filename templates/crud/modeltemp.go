package crudTemplate

import (
	"fmt"
	"strings"
)

func CrudModelTemplate(moduleName string) string {
	return fmt.Sprintf(
		`package models



type %[1]s struct {
}


`, moduleName, strings.ToLower(moduleName))
}
