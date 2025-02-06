package config

import "github.com/spf13/viper"

type Config struct {
	Production_Url    string `mapstructure:"PRODUCTION_URL"`
	Development_Url   string `mapstructure:"DEVELOPMENT_URL"`
	Port              string `mapstructure:"PORT"`
	Customer_Svc_Url  string `mapstructure:"CUSTOMER_SVC_URL"`
	Customer_Svc_Port string `mapstructure:"CUSTOMER_SVC_PORT"`
	Whatsapp_Svc_Url  string `mapstructure:"WHATSAPP_SVC_URL"`
	Whatsapp_Svc_Port string `mapstructure:"WHATSAPP_SVC_PORT"`
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
