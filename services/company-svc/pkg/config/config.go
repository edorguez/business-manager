package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	CompanySvcPort             string `mapstructure:"COMPANY_SVC_PORT"`
	PostgresDBDriver           string `mapstructure:"POSTGRES_DB_DRIVER"`
	CompanyDBSourceDevelopment string `mapstructure:"COMPANY_DB_SOURCE_DEVELOPMENT"`
	CompanyDBSourceProduction  string `mapstructure:"COMPANY_DB_SOURCE_PRODUCTION"`
	FileSvcUrl                 string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort                string `mapstructure:"FILE_SVC_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using .env file for configuration")
	} else {
		log.Println("No .env file found, using environment variables")
	}

	err = viper.Unmarshal(&config)
	return
}
