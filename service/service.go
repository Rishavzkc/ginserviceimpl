package service

import (
	"fmt"
	"net/http"

	"github.com/Rishavzkc/ginserviceimpl/controller"
	"github.com/Rishavzkc/ginserviceimpl/model"
	"github.com/gin-gonic/gin"
)

type CompanyServer interface {
	CreateCompany(c *gin.Context)
	GetAllCompanies(c *gin.Context)
	GetCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}

type CompanyService struct {
	database controller.DatabaseController
}

func NewCompanyService(db controller.DatabaseController) CompanyServer {
	return &CompanyService{
		database: db,
	}
}

func (cs *CompanyService) CreateCompany(c *gin.Context) {
	var company model.Company
	if err := c.BindJSON(&company); err != nil {
		serviceErr := DBInsertionFailure
		serviceErr.Error = fmt.Sprintf("failed to parse record from input: %v", err)
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	if company.ID == "" {
		serviceErr := MissingMandatoryFields
		serviceErr.Error = fmt.Sprintf("missing mandatory field: 'id'")
		return
	}

	if err := cs.database.CreateCompany(&company); err != nil {
		serviceErr := DBInsertionFailure
		serviceErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	resp := SuccessStatus
	resp.Message = "Companies database creation successful"
	c.JSON(http.StatusOK, resp)
}

func (cs *CompanyService) GetAllCompanies(c *gin.Context) {
	companies, err := cs.database.GetAllCompanies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DBRetrievalFailure)
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (cs *CompanyService) GetCompany(c *gin.Context) {
	id := c.Param("id")
	company, err := cs.database.GetCompanyByID(id)
	if err != nil {
		serviceErr := DBRetrievalFailure
		serviceErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}
	c.JSON(http.StatusOK, company)
}

func (cs *CompanyService) UpdateCompany(c *gin.Context) {
	var company model.Company
	if err := c.BindJSON(&company); err != nil {
		serviceErr := DBUpdateFailure
		serviceErr.Error = fmt.Sprintf("failed to parse record from input: %v", err)
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	if company.ID == "" {
		serviceErr := MissingMandatoryFields
		serviceErr.Error = fmt.Sprintf("missing mandatory field: 'id'")
		return
	}

	if err := cs.database.UpdateCompany(&company); err != nil {
		serviceErr := DBUpdateFailure
		serviceErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	resp := SuccessStatus
	resp.Message = fmt.Sprintf("Record with id '%s' updated successfully", company.ID)
	c.JSON(http.StatusOK, resp)
}

func (cs *CompanyService) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		serviceErr := MissingMandatoryFields
		serviceErr.Error = fmt.Sprintf("missing mandatory field: 'id'")
		return
	}

	err := cs.database.DeleteCompanyByID(id)
	if err != nil {
		serviceErr := DBDeleteFailure
		serviceErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, serviceErr)
		return
	}

	successMessage := SuccessStatus
	successMessage.Message = fmt.Sprintf("Successfully deleted record with id '%s'", id)
	c.JSON(http.StatusOK, successMessage)
}
