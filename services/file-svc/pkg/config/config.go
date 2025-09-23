package config

import "github.com/spf13/viper"

type Config struct {
	FileSvcPort          string `mapstructure:"FILE_SVC_PORT"`
	AwsAccessKeyId       string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKeyId string `mapstructure:"AWS_SECRET_ACCESS_KEY_ID"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
