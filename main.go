package main

import (
	"Interfaceginrestapi/config"
	"Interfaceginrestapi/controllers"
	"Interfaceginrestapi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                      = config.SetupDatabaseConnection()
	companyService    services.CompanyService       = services.NewCompanyService(db)
	companyController controllers.CompanyController = controllers.New(companyService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	bookRoutes := r.Group("api/company")
	{
		bookRoutes.GET("/", companyController.GetAll)
		bookRoutes.POST("/", companyController.CreateCompany)
		bookRoutes.GET("/:id", companyController.GetCompany)
		bookRoutes.PUT("/:id", companyController.UpdateCompany)
		bookRoutes.DELETE("/:id", companyController.DeleteCompany)
	}

	r.Run()
}
