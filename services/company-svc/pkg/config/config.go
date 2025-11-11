package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	CompanySvcPort                 string `mapstructure:"COMPANY_SVC_PORT"`
	PostgresDBDriver               string `mapstructure:"POSTGRES_DB_DRIVER"`
	CompanyDBSourceDevelopment     string `mapstructure:"COMPANY_DB_SOURCE_DEVELOPMENT"`
	CompanyDBSourceDockerContainer string `mapstructure:"COMPANY_DB_SOURCE_DOCKER_CONTAINER"`
	FileSvcUrl                     string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort                    string `mapstructure:"FILE_SVC_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using .env file for configuration")
	} else {
		log.Println("No .env file found, using environment variables")

		config.CompanySvcPort = getViperString("COMPANY_SVC_PORT")
		config.PostgresDBDriver = getViperString("POSTGRES_DB_DRIVER")
		config.CompanyDBSourceDevelopment = getViperString("COMPANY_DB_SOURCE_DEVELOPMENT")
		config.CompanyDBSourceDockerContainer = getViperString("COMPANY_DB_SOURCE_DOCKER_CONTAINER")
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
