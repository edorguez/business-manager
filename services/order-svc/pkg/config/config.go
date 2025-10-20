package config

import (
	"log"

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
	}

	err = viper.Unmarshal(&config)
	return
}
