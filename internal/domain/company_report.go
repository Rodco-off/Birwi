package domain

import "time"

type CompanyReport struct {
	ID                        uint      `gorm:"PrimaryKey; autoIncrement"`
	CompanyID                 uint      `gorm:"not null; index"`
	DateReport                time.Time `gorm:"type: date; not null; index"`
	TotalNonCurrentAssets     float64   `gorm:"type: decimal(20,2); not null; index"`
	TotalCurrentAssets        float64   `gorm:"type: decimal(20,2); not null; index"`
	TotalLongTermLiabilities  float64   `gorm:"type: decimal(20,2); not null; index"`
	TotalShortTermLiabilities float64   `gorm:"type: decimal(20,2); not null; index"`
	TotalLiabilitiesAndEquity float64   `gorm:"type: decimal(20,2); not null; index"`
	Revenue                   float64   `gorm:"type: decimal(20,2); not null; index"`
	OperationalProfit         float64   `gorm:"type: decimal(20,2); not null; index"`
	ProfitBeforeTax           float64   `gorm:"type: decimal(20,2); not null; index"`
	NetProfit                 float64   `gorm:"type: decimal(20,2); not null; index"`
}
