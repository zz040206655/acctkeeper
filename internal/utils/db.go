package utils

import (
	"acctkeeper/internal/config"
	"acctkeeper/internal/model"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: refine the connection setting
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetConnMaxLifetime(5 * time.Minute)

	if err = DB.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&model.Account{}, &model.Transaction{}, &model.Report{})
	DB.Model(&model.Transaction{}).AddUniqueIndex("account_id", "amount", "type", "tx_time")
	fmt.Println("Success connecting to MySQL")
}
