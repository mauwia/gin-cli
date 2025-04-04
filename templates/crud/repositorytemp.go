package crudTemplate

import (
	"fmt"
	"strings"
)

func CrudRepositoryTemplate(moduleName string, projectName string) string {
	return fmt.Sprintf(
		`package repositories

import (

    "%[3]s/internal/models"

)

type %[1]sRepository struct {
}

func New%[1]sRepository() *%[1]sRepository {
    return &%[1]sRepository{}
}

func (s *%[1]sRepository) Create%[1]s(%[2]s *models.%[1]s) error {
    return nil 
}

func (s *%[1]sRepository) Get%[1]sByID(id uint) (*models.%[1]s, error) {
    return nil,nil
}
func (s *%[1]sRepository) GetAll%[1]ss() ([]*models.%[1]s, error) {
    return nil,nil
}
func (s *%[1]sRepository) Update%[1]s(%[2]s *models.%[1]s,id uint) (*models.%[1]s, error) {
	return nil,nil
}
func (s *%[1]sRepository) Delete%[1]s(id uint) error {
    return nil
}
`, moduleName, strings.ToLower(moduleName), projectName)
}
