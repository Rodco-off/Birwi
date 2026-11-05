package repository

import (
	"github.com/Rodco-off/Birwi/internal/domain"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (repo *CompanyRepository) Create(sub *domain.Company) error {
	return repo.db.Create(sub).Error
}

func (repo *CompanyRepository) Update(sub *domain.Company) error {
	return repo.db.Save(sub).Error
}

func (repo *CompanyRepository) GetByID(id uint) (*domain.Company, error) {
	var company domain.Company
	if err := repo.db.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (repo *CompanyRepository) GetAll() ([]domain.Company, error) {
	var companies []domain.Company
	if err := repo.db.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (repo *CompanyRepository) GetByCompanyType(sub *domain.Company) error {
	var companies []domain.Company
	if err := repo.db.Where("company_type = ?", sub.CompanyType).Find(&companies).Error; err != nil {
		return err
	}
	return nil
}
