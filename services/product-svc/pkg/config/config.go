package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ProductSvcPort             string `mapstructure:"PRODUCT_SVC_PORT"`
	ProductDBName              string `mapstructure:"PRODUCT_DB_NAME"`
	ProductDBSourceDevelopment string `mapstructure:"PRODUCT_DB_SOURCE_DEVELOPMENT"`
	ProductDBSourceProduction  string `mapstructure:"PRODUCT_DB_SOURCE_PRODUCTION"`
	CompanySvcUrl              string `mapstructure:"COMPANY_SVC_URL"`
	CompanySvcPort             string `mapstructure:"COMPANY_SVC_PORT"`
	FileSvcUrl                 string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort                string `mapstructure:"FILE_SVC_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using .env file for configuration")
	} else {
		log.Println("No .env file found, using environment variables")

		config.ProductSvcPort = getViperString("PRODUCT_SVC_PORT")
		config.ProductDBName = getViperString("PRODUCT_DB_NAME")
		config.ProductDBSourceDevelopment = getViperString("PRODUCT_DB_SOURCE_DEVELOPMENT")
		config.ProductDBSourceProduction = getViperString("PRODUCT_DB_SOURCE_PRODUCTION")
		config.CompanySvcUrl = getViperString("COMPANY_SVC_URL")
		config.CompanySvcPort = getViperString("COMPANY_SVC_PORT")
		config.FileSvcUrl = getViperString("FILE_SVC_URL")
		config.FileSvcPort = getViperString("FILE_SVC_PORT")
	}

	err = viper.Unmarshal(&config)
	return
}

func getViperString(key string) string {
	// First try to get from viper (which includes environment variables)
	if value := viper.GetString(key); value != "" {
		return value
	}
	// Fallback to direct environment variable
	return os.Getenv(key)
}
