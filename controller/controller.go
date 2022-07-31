package controller

import (
	"fmt"

	"github.com/Rishavzkc/ginserviceimpl/model"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type DatabaseController interface {
	CreateCompany(company *model.Company) error
	GetAllCompanies(ctx *gin.Context) ([]model.Company, error)
	GetCompanyByID(id string) (*model.Company, error)
	UpdateCompany(company *model.Company) error
	DeleteCompanyByID(id string) error
}

type CompanyController struct {
	companyDB *gorm.DB
}

func New(db *gorm.DB) *CompanyController {
	return &CompanyController{
		companyDB: db,
	}
}

func (c *CompanyController) CreateCompany(company *model.Company) error {

	if tx := c.companyDB.Create(company); tx.Error != nil {
		return fmt.Errorf("failed to insert record in company db: %w", tx.Error)
	}
	return nil
}

func (c *CompanyController) GetAllCompanies(ctx *gin.Context) ([]model.Company, error) {
	var companies []model.Company
	if tx := c.companyDB.Find(&companies); tx.Error != nil {
		return nil, fmt.Errorf("failed to fetch all records from company db: %w", tx.Error)
	}
	return companies, nil
}

func (c *CompanyController) GetCompanyByID(id string) (*model.Company, error) {
	var company model.Company
	if tx := c.companyDB.First(&company, "id = ?", id); tx.Error != nil {
		return nil, fmt.Errorf("failed to fetch record with id '%s': %w", id, tx.Error)
	}
	return &company, nil
}

func (c *CompanyController) UpdateCompany(company *model.Company) error {
	tx := c.companyDB.Model(model.Company{ID: company.ID}).Updates(&model.Company{
		Name:     company.Name,
		Location: company.Location,
	})
	if tx.Error != nil {
		return fmt.Errorf("failed to update record with id '%s' in company db: %w", company.ID, tx.Error)
	}
	return nil
}

func (c *CompanyController) DeleteCompanyByID(id string) error {
	tx := c.companyDB.Delete(&model.Company{ID: id})
	if tx.Error != nil {
		return fmt.Errorf("failed to delete record with id '%s': %w", id, tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("no record found to delete for id '%s'", id)
	}
	return nil
}
