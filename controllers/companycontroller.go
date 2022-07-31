package controllers

import (
	"Interfaceginrestapi/models"
	"Interfaceginrestapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	CompanyService services.CompanyService
}

func New(companyservice services.CompanyService) CompanyController {
	return CompanyController{
		CompanyService: companyservice,
	}
}

func (cc *CompanyController) CreateCompany(ctx *gin.Context) {
	var company models.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CompanyService.CreateCompany(&company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (cc *CompanyController) GetCompany(ctx *gin.Context) {
	companyname := ctx.Param("name")
	company, err := cc.CompanyService.GetCompany(&companyname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//ctx.JSON(http.StatusOK, company)
	ctx.JSON(http.StatusOK, gin.H{"data": company})

}

func (cc *CompanyController) GetAll(ctx *gin.Context) {
	companies, err := cc.CompanyService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func (cc *CompanyController) UpdateCompany(ctx *gin.Context) {
	var company models.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CompanyService.UpdateCompany(&company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	//ctx.JSON(http.StatusOK, gin.H{"data": company})
}

func (cc *CompanyController) DeleteCompany(ctx *gin.Context) {
	companyname := ctx.Param("name")
	err := cc.CompanyService.DeleteCompany(&companyname)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}
