package main

import (
	"fmt"
	"log"

	"github.com/Rishavzkc/ginserviceimpl/config"
	"github.com/Rishavzkc/ginserviceimpl/controller"
	"github.com/Rishavzkc/ginserviceimpl/database"
	"github.com/Rishavzkc/ginserviceimpl/service"
	"github.com/gin-gonic/gin"
)

func main() {

	configs := config.NewConfig()
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.Database,
	)

	db, err := database.SetupDatabaseConnection(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := database.CloseDatabaseConnection(db); err != nil {
			log.Fatal(err)
		}
	}()

	r := gin.Default()

	dbController := controller.New(db)
	companyService := service.NewCompanyService(dbController)

	companyRouter := r.Group("/company")
	{
		companyRouter.POST("/", companyService.CreateCompany)
		companyRouter.GET("/", companyService.GetAllCompanies)
		companyRouter.GET("/:id", companyService.GetCompany)
		companyRouter.PUT("/", companyService.UpdateCompany)
		companyRouter.DELETE("/:id", companyService.DeleteCompany)
	}

	if err := r.Run(configs.ServiceHost); err != nil {
		log.Fatal("failure at running server: %w", err)
	}
}
