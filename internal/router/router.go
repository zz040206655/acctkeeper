package router

import (
	"acctkeeper/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", handler.Register)
	r.POST("/transaction", handler.AddTransaction)
	r.POST("/import_transactions", handler.ImportTransactions)
	r.GET("/:username/report", handler.GetReport)
}
