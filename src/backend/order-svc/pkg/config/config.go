package config

import "github.com/spf13/viper"

type Config struct {
	Environment      string `mapstructure:"ENVIRONMENT"`
	Port             string `mapstructure:"PORT"`
	Customer_Svc_Url string `mapstructure:"CUSTOMER_SVC_URL"`
	Whatsapp_Svc_Url string `mapstructure:"WHATSAPP_SVC_URL"`
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
