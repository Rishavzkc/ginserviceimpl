package services

import (
	"Interfaceginrestapi/config"
	"Interfaceginrestapi/models"
	"database/sql"

	"gorm.io/gorm"
)

var conn *sql.DB

type CompanyServiceImpl struct {
	companycollection *gorm.DB
}

func NewCompanyService(companycollection *gorm.DB) CompanyService {
	return &CompanyServiceImpl{
		companycollection: companycollection,
	}
}

func (c *CompanyServiceImpl) CreateCompany(company *models.Company) error {

	config.SetupDatabaseConnection().Create(&company)
	return nil
}

func (c *CompanyServiceImpl) GetCompany(name *string) (*models.Company, error) {
	var company models.Company

	c.companycollection.Where("name=?", name).First(name)

	c.companycollection.Preload("Company").Find(&company)
	return &company, nil

}

func (c *CompanyServiceImpl) GetAll() ([]*models.Company, error) {
	var companies []*models.Company
	c.companycollection.Find(&companies)
	return companies, nil

}

func (c *CompanyServiceImpl) UpdateCompany(company *models.Company) error {
	c.companycollection.Where("name =?", company.Name).First(&company)
	c.companycollection.Preload("Company").Find(&company)
	c.companycollection.Updates(company)
	return nil
}

func (c *CompanyServiceImpl) DeleteCompany(name *string) error {
	var company models.Company
	c.companycollection.Delete(company)
	return nil
}
