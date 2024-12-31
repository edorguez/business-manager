package config

import "github.com/spf13/viper"

type Config struct {
	Production_Url           string `mapstructure:"PRODUCTION_URL"`
	Development_Url          string `mapstructure:"DEVELOPMENT_URL"`
	Port                     string `mapstructure:"PORT"`
	Aws_Access_Key_Id        string `mapstructure:"AWS_ACCESS_KEY_ID"`
	Aws_Secret_Access_Key_Id string `mapstructure:"AWS_SECRET_ACCESS_KEY_ID"`
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
