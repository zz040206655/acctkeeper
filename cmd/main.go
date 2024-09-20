package main

import (
	"acctkeeper/internal/config"
	"acctkeeper/internal/router"
	"acctkeeper/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.InitConfig()
	utils.InitDB()

	defer utils.DB.Close()

	r := gin.Default()

	router.RegisterRoutes(r)

	if err := r.Run(config.ServerPort); err != nil {
		log.Fatalf("Failed to start gin service: %v\n", err)
	}

}
