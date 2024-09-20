package config

import (
	"github.com/spf13/viper"
)

var (
	ServerPort string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	ServerPort = viper.GetString("server.port")
	DBUser = viper.GetString("database.user")
	DBPassword = viper.GetString("database.password")
	DBHost = viper.GetString("database.host")
	DBPort = viper.GetString("database.port")
	DBName = viper.GetString("database.name")
}
