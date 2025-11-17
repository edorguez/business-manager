package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AuthSvcPort                 string `mapstructure:"AUTH_SVC_PORT"`
	PostgresDBDriver            string `mapstructure:"POSTGRES_DB_DRIVER"`
	AuthDBSourceDevelopment     string `mapstructure:"AUTH_DB_SOURCE_DEVELOPMENT"`
	AuthDBSourceDockerContainer string `mapstructure:"AUTH_DB_SOURCE_DOCKER_CONTAINER"`
	JWTSecretKey                string `mapstructure:"JWT_SECRET_KEY"`
	CompanySvcUrl               string `mapstructure:"COMPANY_SVC_URL"`
	CompanySvcPort              string `mapstructure:"COMPANY_SVC_PORT"`
	WhatsappSvcUrl              string `mapstructure:"WHATSAPP_SVC_URL"`
	WhatsappSvcPort             string `mapstructure:"WHATSAPP_SVC_PORT"`
	FileSvcUrl                  string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort                 string `mapstructure:"FILE_SVC_PORT"`
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

		config.AuthSvcPort = getViperString("AUTH_SVC_PORT")
		config.PostgresDBDriver = getViperString("POSTGRES_DB_DRIVER")
		config.AuthDBSourceDevelopment = getViperString("AUTH_DB_SOURCE_DEVELOPMENT")
		config.AuthDBSourceDockerContainer = getViperString("AUTH_DB_SOURCE_DOCKER_CONTAINER")
		config.JWTSecretKey = getViperString("JWT_SECRET_KEY")
		config.CompanySvcUrl = getViperString("COMPANY_SVC_URL")
		config.CompanySvcPort = getViperString("COMPANY_SVC_PORT")
		config.WhatsappSvcUrl = getViperString("WHATSAPP_SVC_URL")
		config.WhatsappSvcPort = getViperString("WHATSAPP_SVC_PORT")
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
