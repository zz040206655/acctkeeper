package handler

import (
	"acctkeeper/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetReport(c *gin.Context) {
	username := c.Param("username")
	year, month := c.Query("year"), c.Query("month")

	report, err := service.GetReport(username, year, month)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
