package crudTemplate

import (
	"fmt"
	"strings"
)

func CrudControllerTemplate(moduleName string, projectName string) string {
	return fmt.Sprintf(
		`package controllers

import (
    "%[3]s/internal/services"
    "%[3]s/internal/models"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type %[1]sController struct {
    Service *services.%[1]sService
}

func New%[1]sController(service *services.%[1]sService) *%[1]sController {
    return &%[1]sController{Service: service}
}

func (ctrl *%[1]sController) Create%[1]s(c *gin.Context)  {
    var %[2]s models.%[1]s
    if err := c.ShouldBindJSON(&%[2]s); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.Service.Create%[1]s(&%[2]s); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, %[2]s)

}

func (ctrl *%[1]sController) Get%[1]sByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    %[2]s, err := ctrl.Service.Get%[1]sByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "%[1]s not found"})
        return
    }
    c.JSON(http.StatusOK, %[2]s)
}
func (ctrl *%[1]sController) GetAll%[1]ss(c *gin.Context) {
    %[2]ss, err := ctrl.Service.GetAll%[1]ss()
	if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "%[1]ss not found"})
        return
    }
    c.JSON(http.StatusOK, %[2]ss)
}
func (ctrl *%[1]sController) Update%[1]s(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
	var %[2]s models.%[1]s
    if err := c.ShouldBindJSON(&%[2]s); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updated%[2]s, err := ctrl.Service.Update%[1]s(&%[2]s,uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "%[1]s not found"})
        return
    }
    c.JSON(http.StatusOK, updated%[2]s)
}
func (ctrl *%[1]sController) Delete%[1]s(c *gin.Context) {
      id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    err = ctrl.Service.Delete%[1]s(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "%[1]s not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
`, moduleName, strings.ToLower(moduleName), projectName)
}
