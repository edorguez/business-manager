package config

import (
	"log"

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
	}

	err = viper.Unmarshal(&config)
	return
}
