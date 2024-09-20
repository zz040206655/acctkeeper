package service

import (
	"acctkeeper/internal/model"
	"acctkeeper/internal/utils"
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetReport(username, year, month string) (model.Report, error) {
	var report model.Report
	var account model.Account
	if err := utils.DB.Where("username = ?", username).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return report, errors.New("account not found")
		}
		return report, err
	}

	sigma := "account_id = ? AND year = ? AND month = ?"
	err := utils.DB.Where(sigma, account.ID, year, month).First(&report).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return GenReport(account, year, month)
		}
		return report, err
	}
	return report, nil
}

func GenReport(account model.Account, year, month string) (model.Report, error) {
	var report model.Report
	var transactions []model.Transaction

	sigma := "account_id = ? AND YEAR(tx_time) = ? AND MONTH(tx_time) = ?"
	err := utils.DB.Where(sigma, account.ID, year, month).Find(&transactions).Error
	if err != nil {
		return report, err
	}

	var totalIncome, totalExpense float64
	for _, tx := range transactions {
		if tx.Amount > 0 {
			totalIncome += tx.Amount
		} else {
			totalExpense -= tx.Amount
		}
	}

	y, _ := strconv.Atoi(year)
	m, _ := strconv.Atoi(month)

	report = model.Report{
		AccountID:    account.ID,
		Year:         y,
		Month:        m,
		TotalIncome:  totalIncome,
		TotalExpense: totalExpense,
	}

	if err := utils.DB.Create(&report).Error; err != nil {
		return report, err
	}
	return report, nil
}
