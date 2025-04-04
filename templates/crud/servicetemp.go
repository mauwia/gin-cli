package crudTemplate

import (
	"fmt"
	"strings"
)

func CrudServiceTemplate(moduleName string, projectName string) string {
	return fmt.Sprintf(
		`package services

import (
    "%[3]s/internal/repositories"
    "%[3]s/internal/models"
)

type %[1]sService struct {
    Repository *repositories.%[1]sRepository
}

func New%[1]sService(repo *repositories.%[1]sRepository) *%[1]sService {
    return &%[1]sService{Repository: repo}
}

func (s *%[1]sService) Create%[1]s(%[2]s *models.%[1]s) error {
    return s.Repository.Create%[1]s(%[2]s)
}

func (s *%[1]sService) Get%[1]sByID(id uint) (*models.%[1]s, error) {
    return s.Repository.Get%[1]sByID(id)
}
func (s *%[1]sService) GetAll%[1]ss() ([]*models.%[1]s, error) {
    return s.Repository.GetAll%[1]ss()
}
func (s *%[1]sService) Update%[1]s(%[2]s *models.%[1]s,id uint) (*models.%[1]s, error) {
	return s.Repository.Update%[1]s(%[2]s,id)
}
func (s *%[1]sService) Delete%[1]s(id uint) error {
    return s.Repository.Delete%[1]s(id)
}
`, moduleName, strings.ToLower(moduleName), projectName)
}
