package service

import (
	"acctkeeper/internal/model"
	"acctkeeper/internal/utils"
	"errors"

	"github.com/jinzhu/gorm"
)

func AddTransaction(txReq model.TransactionReq) (float64, error) {
	var tx model.Transaction
	var account model.Account
	if err := utils.DB.Where("username = ?", txReq.Username).First(&account).Error; err != nil {
		return 0.0, errors.New("account not found")
	}

	// Check and skip the same transaction.
	sigma := "account_id = ? AND amount = ? AND type = ? AND tx_time = ?"
	err := utils.DB.Where(sigma, account.ID, txReq.Amount, txReq.Type, txReq.TxTime).First(&tx).Error
	if err == nil {
		return 0.0, errors.New("got the same transaction")
	}

	tx = model.Transaction{
		AccountID: account.ID,
		Amount:    txReq.Amount,
		Type:      txReq.Type,
		TxTime:    txReq.TxTime,
	}

	err = utils.DB.Create(&tx).Error
	if err != nil {
		return 0.0, err
	}

	// Update latest balance
	account.Balance += txReq.Amount
	err = utils.DB.Save(&account).Error
	if err != nil {
		return 0.0, err
	}

	return account.Balance, nil
}

func ImportTransactions(txReqs []model.TransactionReq) (float64, error) {
	if len(txReqs) == 0 {
		return 0.0, errors.New("no transactions in this batch")
	}

	var account model.Account
	if err := utils.DB.Where("username = ?", txReqs[0].Username).First(&account).Error; err != nil {
		return 0.0, errors.New("account not found")
	}

	newtxs := []model.Transaction{}
	for _, r := range txReqs {
		var tx model.Transaction

		// Check and skip the same transaction.
		sigma := "account_id = ? AND amount = ? AND type = ? AND tx_time = ?"
		err := utils.DB.Where(sigma, account.ID, r.Amount, r.Type, r.TxTime).First(&tx).Error
		if err == nil {
			continue
		} else if err != gorm.ErrRecordNotFound {
			return 0.0, err
		}

		account.Balance += tx.Amount
		newtxs = append(newtxs, tx)
	}

	if len(newtxs) > 0 {
		if err := utils.DB.Create(&newtxs).Error; err != nil {
			return 0.0, err
		}
	}

	if err := utils.DB.Save(&account).Error; err != nil {
		return 0.0, err
	}

	return account.Balance, nil
}
