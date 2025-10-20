package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	CustomerSvcPort             string `mapstructure:"CUSTOMER_SVC_PORT"`
	PostgresDBDriver            string `mapstructure:"POSTGRES_DB_DRIVER"`
	CustomerDBSourceDevelopment string `mapstructure:"CUSTOMER_DB_SOURCE_DEVELOPMENT"`
	CustomerDBSourceProduction  string `mapstructure:"CUSTOMER_DB_SOURCE_PRODUCTION"`
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

		config.CustomerSvcPort = getViperString("CUSTOMER_SVC_PORT")
		config.PostgresDBDriver = getViperString("POSTGRES_DB_DRIVER")
		config.CustomerDBSourceDevelopment = getViperString("CUSTOMER_DB_SOURCE_DEVELOPMENT")
		config.CustomerDBSourceProduction = getViperString("CUSTOMER_DB_SOURCE_PRODUCTION")
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
