package model

import (
	"time"
)

/*
Transaction Request Format:
{
    "account_id": 1,
    "amount": -50.0,
    "type": "cash",
    "time": "2024-09-01T15:04:05Z"
}
*/

type TransactionReq struct {
	Username string    `json:"username" binding:"required"`
	Amount   float64   `json:"amount" binding:"required"`
	Type     string    `json:"type" binding:"required"`
	TxTime   time.Time `json:"txtime" binding:"required"`
}

type Transaction struct {
	ID        uint      `gorm:"primaryKey"`
	AccountID uint      `gorm:"not null"` // connect to Account ID
	Amount    float64   `gorm:"not null"`
	Type      string    `gorm:"size:50;not null"` // cash or bank
	TxTime    time.Time `gorm:"not null"`
}
