package service

import (
	"acctkeeper/internal/model"
	"acctkeeper/internal/utils"
	"errors"
)

func Register(username string) error {
	var account model.Account

	if err := utils.DB.Where("username = ?", username).First(&account).Error; err == nil {
		return errors.New("username already exist")
	}

	account = model.Account{
		Username: username,
		Balance:  0.0,
	}

	if err := utils.DB.Create(&account).Error; err != nil {
		return err
	}

	return nil
}
