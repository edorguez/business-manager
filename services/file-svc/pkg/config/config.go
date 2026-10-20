package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	FileSvcPort          string `mapstructure:"FILE_SVC_PORT"`
	AwsAccessKeyId       string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKeyId string `mapstructure:"AWS_SECRET_ACCESS_KEY_ID"`

	// S3-compatible storage (VersityGW/MinIO in self-hosted deployments).
	S3Endpoint       string `mapstructure:"S3_ENDPOINT"`
	S3Region         string `mapstructure:"S3_REGION"`
	S3PublicUrl      string `mapstructure:"S3_PUBLIC_URL"`
	S3ForcePathStyle bool   `mapstructure:"S3_FORCE_PATH_STYLE"`
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
		config.S3Endpoint = getViperString("S3_ENDPOINT")
		config.S3Region = getViperString("S3_REGION")
		config.S3PublicUrl = getViperString("S3_PUBLIC_URL")
		config.S3ForcePathStyle = strings.EqualFold(getViperString("S3_FORCE_PATH_STYLE"), "true")
	}

	err = viper.Unmarshal(&config)

	if config.S3Region == "" {
		config.S3Region = "us-east-1"
	}
	if config.S3PublicUrl == "" {
		config.S3PublicUrl = "/media"
	}
	if forcePathStyle := getViperString("S3_FORCE_PATH_STYLE"); forcePathStyle == "" {
		config.S3ForcePathStyle = true
	} else {
		config.S3ForcePathStyle = strings.EqualFold(forcePathStyle, "true")
	}

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
