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

	companyRoutes := r.Group("/company")
	{
		companyRoutes.GET("/", companyController.GetAll)
		companyRoutes.POST("/", companyController.CreateCompany)
		companyRoutes.GET("/:id", companyController.GetCompany)
		companyRoutes.PUT("/:id", companyController.UpdateCompany)
		companyRoutes.DELETE("/:id", companyController.DeleteCompany)
	}

	r.Run()
}
