package db

import (
	"fmt"

	"github.com/Rodco-off/Birwi/internal/config"
	"github.com/Rodco-off/Birwi/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.UserDB, cfg.UserDB, cfg.PasswordDB, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, err
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Company{},
		&domain.CompanyReport{},
	)
}
