package crudTemplate

import (
	"fmt"
	"strings"
)

func CrudRouterTemplate(moduleName string, projectName string) string {
	return fmt.Sprintf(`

func Setup%[1]sRoutes(r *gin.Engine, cfg *config.AppConfig) {

    %[2]sRepo:= repositories.New%[1]sRepository()
    %[2]sService:= services.New%[1]sService(%[2]sRepo)
	%[2]sController:=controllers.New%[1]sController(%[2]sService)

	%[2]s:= r.Group("/api/%[2]s") 
	{
		%[2]s.POST("/", %[2]sController.Create%[1]s)
		%[2]s.GET("/all", %[2]sController.GetAll%[1]ss)
		%[2]s.GET("/:id", %[2]sController.Get%[1]sByID)
		%[2]s.PATCH("/:id", %[2]sController.Update%[1]s)
		%[2]s.DELETE("/:id", %[2]sController.Delete%[1]s)
	}

}
    `, moduleName, strings.ToLower(moduleName))

}
