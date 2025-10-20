package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	FileSvcPort          string `mapstructure:"FILE_SVC_PORT"`
	AwsAccessKeyId       string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKeyId string `mapstructure:"AWS_SECRET_ACCESS_KEY_ID"`
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

		config.FileSvcPort = getViperString("FILE_SVC_PORT")
		config.AwsAccessKeyId = getViperString("AWS_ACCESS_KEY_ID")
		config.AwsSecretAccessKeyId = getViperString("AWS_SECRET_ACCESS_KEY_ID")
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
