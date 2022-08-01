package services

import (
	"Interfaceginrestapi/models"

	"fmt"

	"gorm.io/gorm"
)

type CompanyServiceImpl struct {
	companycollection *gorm.DB
}

func NewCompanyService(companycollection *gorm.DB) CompanyService {
	return &CompanyServiceImpl{
		companycollection: companycollection,
	}
}

func (c *CompanyServiceImpl) CreateCompany(company *models.Company) error {
	if tx := c.companycollection.Create(company); tx.Error != nil {
		return fmt.Errorf("failed to insert record in company db: %w", tx.Error)
	}
	return nil

}

func (c *CompanyServiceImpl) GetCompany(id string) (*models.Company, error) {

	var company models.Company
	if tx := c.companycollection.First(&company, "id = ?", id); tx.Error != nil {
		return nil, fmt.Errorf("failed to fetch record with id '%s': %w", id, tx.Error)
	}
	return &company, nil

}

func (c *CompanyServiceImpl) GetAll() ([]*models.Company, error) {

	var companies []*models.Company
	if tx := c.companycollection.Find(&companies); tx.Error != nil {
		return nil, fmt.Errorf("failed to fetch all records from company db: %w", tx.Error)
	}
	return companies, nil

}

func (c *CompanyServiceImpl) UpdateCompany(company *models.Company) error {

	tx := c.companycollection.Model(models.Company{Id: company.Id}).Updates(&models.Company{
		Name:     company.Name,
		Location: company.Location,
	})
	if tx.Error != nil {
		return fmt.Errorf("failed to update record with id '%s' in company db: %w", company.Id, tx.Error)
	}
	return nil
}

func (c *CompanyServiceImpl) DeleteCompany(id string) error {
	tx := c.companycollection.Delete(&models.Company{Id: id})
	if tx.Error != nil {
		return fmt.Errorf("failed to delete record with id '%s': %w", id, tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("no record found to delete for id '%s'", id)
	}
	return nil

}
