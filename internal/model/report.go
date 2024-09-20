package model

import "time"

type Report struct {
	ID           uint    `gorm:"primaryKey"`
	AccountID    uint    `gorm:"not null"`
	Year         int     `gorm:"not null;uniqueIndex:idx_account_year_month"`
	Month        int     `gorm:"not null;uniqueIndex:idx_account_year_month"`
	TotalIncome  float64 `gorm:"not null"`
	TotalExpense float64 `gorm:"not null"`
	CreatedAt    time.Time
}
