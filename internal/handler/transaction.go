package handler

import (
	"acctkeeper/internal/model"
	"acctkeeper/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTransaction(c *gin.Context) {
	var txReq model.TransactionReq

	if err := c.ShouldBindJSON(&txReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	balance, err := service.AddTransaction(txReq)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	msg := fmt.Sprintf("Name: %s, current balance: %f", txReq.Username, balance)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func ImportTransactions(c *gin.Context) {
	var txReqs []model.TransactionReq

	if err := c.ShouldBindJSON(&txReqs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, err := service.ImportTransactions(txReqs)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf("Name: %s, current balance: %f", txReqs[0].Username, balance)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
