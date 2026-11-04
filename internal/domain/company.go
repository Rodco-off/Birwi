package domain

import "time"

type Company struct {
	ID           uint      `gorm:"PrimaryKey; index; autoIncrement"`
	Name         string    `gorm:"type: date; not null; index"`
	DateFounded  time.Time `gorm:"type: string; not null"`
	Country      string    `gorm:"type: string; not null"`
	CompanyType  string    `gorm:"type: string; not null"`
	CompanyOwner string    `gorm:"type: string; not null"`
	Address      string    `gorm:"type: string; not null"`
}
