package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	OrderSvcPort    string `mapstructure:"ORDER_SVC_PORT"`
	CustomerSvcUrl  string `mapstructure:"CUSTOMER_SVC_URL"`
	CustomerSvcPort string `mapstructure:"CUSTOMER_SVC_PORT"`
	WhatsappSvcUrl  string `mapstructure:"WHATSAPP_SVC_URL"`
	WhatsappSvcPort string `mapstructure:"WHATSAPP_SVC_PORT"`
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
