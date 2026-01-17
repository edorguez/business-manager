package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	OrderSvcPort                 string `mapstructure:"ORDER_SVC_PORT"`
	CustomerSvcUrl               string `mapstructure:"CUSTOMER_SVC_URL"`
	CustomerSvcPort              string `mapstructure:"CUSTOMER_SVC_PORT"`
	WhatsappSvcUrl               string `mapstructure:"WHATSAPP_SVC_URL"`
	WhatsappSvcPort              string `mapstructure:"WHATSAPP_SVC_PORT"`
	PostgresDBDriver             string `mapstructure:"POSTGRES_DB_DRIVER"`
	OrderDBSourceDevelopment     string `mapstructure:"ORDER_DB_SOURCE_DEVELOPMENT"`
	OrderDBSourceDockerContainer string `mapstructure:"ORDER_DB_SOURCE_DOCKER_CONTAINER"`
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

		config.OrderSvcPort = getViperString("ORDER_SVC_PORT")
		config.CustomerSvcUrl = getViperString("CUSTOMER_SVC_URL")
		config.CustomerSvcPort = getViperString("CUSTOMER_SVC_PORT")
		config.WhatsappSvcUrl = getViperString("WHATSAPP_SVC_URL")
		config.WhatsappSvcPort = getViperString("WHATSAPP_SVC_PORT")
		config.PostgresDBDriver = getViperString("POSTGRES_DB_DRIVER")
		config.OrderDBSourceDevelopment = getViperString("ORDER_DB_SOURCE_DEVELOPMENT")
		config.OrderDBSourceDockerContainer = getViperString("ORDER_DB_SOURCE_DOCKER_CONTAINER")
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
