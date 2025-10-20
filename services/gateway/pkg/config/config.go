package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ProductionUrl   string `mapstructure:"PRODUCTION_URL"`
	DevelopmentUrl  string `mapstructure:"DEVELOPMENT_URL"`
	GatewayPort     string `mapstructure:"GATEWAY_PORT"`
	CustomerSvcUrl  string `mapstructure:"CUSTOMER_SVC_URL"`
	CustomerSvcPort string `mapstructure:"CUSTOMER_SVC_PORT"`
	CompanySvcUrl   string `mapstructure:"COMPANY_SVC_URL"`
	CompanySvcPort  string `mapstructure:"COMPANY_SVC_PORT"`
	ProductSvcUrl   string `mapstructure:"PRODUCT_SVC_URL"`
	ProductSvcPort  string `mapstructure:"PRODUCT_SVC_PORT"`
	AuthSvcUrl      string `mapstructure:"AUTH_SVC_URL"`
	AuthSvcPort     string `mapstructure:"AUTH_SVC_PORT"`
	OrderSvcUrl     string `mapstructure:"ORDER_SVC_URL"`
	OrderSvcPort    string `mapstructure:"ORDER_SVC_PORT"`
	FileSvcUrl      string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort     string `mapstructure:"FILE_SVC_PORT"`
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

		config.ProductionUrl = getViperString("PRODUCTION_URL")
		config.DevelopmentUrl = getViperString("DEVELOPMENT_URL")
		config.GatewayPort = getViperString("GATEWAY_PORT")
		config.CustomerSvcUrl = getViperString("CUSTOMER_SVC_URL")
		config.CustomerSvcPort = getViperString("CUSTOMER_SVC_PORT")
		config.CompanySvcUrl = getViperString("COMPANY_SVC_URL")
		config.CompanySvcPort = getViperString("COMPANY_SVC_PORT")
		config.ProductSvcUrl = getViperString("PRODUCT_SVC_URL")
		config.ProductSvcPort = getViperString("PRODUCT_SVC_PORT")
		config.AuthSvcUrl = getViperString("AUTH_SVC_URL")
		config.AuthSvcPort = getViperString("AUTH_SVC_PORT")
		config.OrderSvcUrl = getViperString("ORDER_SVC_URL")
		config.OrderSvcPort = getViperString("ORDER_SVC_PORT")
		config.FileSvcUrl = getViperString("FILE_SVC_URL")
		config.FileSvcPort = getViperString("FILE_SVC_PORT")
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
