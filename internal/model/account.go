package model

type Account struct {
	ID           uint          `gorm:"primaryKey"`
	Username     string        `gorm:"size:255;not null;unique"`
	Balance      float64       `gorm:"not null;default:0"`
	Transactions []Transaction `gorm:"foreignKey:AccountID"`
	Reports      []Report      `gorm:"foreignKey:AccountID"`
}
