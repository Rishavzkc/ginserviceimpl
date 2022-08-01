package services

import "Interfaceginrestapi/models"

type CompanyService interface {
	CreateCompany(*models.Company) error
	GetCompany(string) (*models.Company, error)
	GetAll() ([]*models.Company, error)
	UpdateCompany(*models.Company) error
	DeleteCompany(string) error
}
