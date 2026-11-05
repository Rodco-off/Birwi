package repository

import (
	"github.com/Rodco-off/Birwi/internal/domain"
	"gorm.io/gorm"
)

type CompanyReportRepository struct {
	db *gorm.DB
}

func NewCompanyReportRepository(db *gorm.DB) *CompanyReportRepository {
	return &CompanyReportRepository{db: db}
}

func (repo *CompanyReportRepository) Create(sub *domain.CompanyReport) error {
	return repo.db.Create(sub).Error
}

func (repo *CompanyReportRepository) Update(sub *domain.CompanyReport) error {
	return repo.db.Save(sub).Error
}

func (repo *CompanyReportRepository) GetByID(id uint) (*domain.CompanyReport, error) {
	var companyReport domain.CompanyReport
	if err := repo.db.First(&companyReport, id).Error; err != nil {
		return nil, err
	}
	return &companyReport, nil
}

func (repo *CompanyReportRepository) GetByCompanyId(companyId uint) ([]domain.CompanyReport, error) {
	var companyReports []domain.CompanyReport
	if err := repo.db.Where("company_id = ?", companyId).Find(&companyReports).Error; err != nil {
		return nil, err
	}
	return companyReports, nil
}
