package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	WhatsappSvcPort                 string `mapstructure:"WHATSAPP_SVC_PORT"`
	WhatsappWsPort                  string `mapstructure:"WHATSAPP_WS_PORT"`
	PostgresDBDriver                string `mapstructure:"POSTGRES_DB_DRIVER"`
	WhatsappDBSourceDevelopment     string `mapstructure:"WHATSAPP_DB_SOURCE_DEVELOPMENT"`
	WhatsappDBSourceDockerContainer string `mapstructure:"WHATSAPP_DB_SOURCE_DOCKER_CONTAINER"`
	TwilioAccountSID                string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken                 string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioPhoneNumber               string `mapstructure:"TWILIO_PHONE_NUMBER"`
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

		config.WhatsappSvcPort = getViperString("WHATSAPP_SVC_PORT")
		config.WhatsappWsPort = getViperString("WHATSAPP_WS_PORT")
		config.PostgresDBDriver = getViperString("POSTGRES_DB_DRIVER")
		config.WhatsappDBSourceDevelopment = getViperString("WHATSAPP_DB_SOURCE_DEVELOPMENT")
		config.WhatsappDBSourceDockerContainer = getViperString("WHATSAPP_DB_SOURCE_DOCKER_CONTAINER")
		config.TwilioAccountSID = getViperString("TWILIO_ACCOUNT_SID")
		config.TwilioAuthToken = getViperString("TWILIO_AUTH_TOKEN")
		config.TwilioPhoneNumber = getViperString("TWILIO_PHONE_NUMBER")
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
