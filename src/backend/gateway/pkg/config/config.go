package config

import "github.com/spf13/viper"

type Config struct {
	Environment      string `mapstructure:"ENVIRONMENT"`
	Gateway_Url      string `mapstructure:"GATEWAY_URL"`
	Customer_Svc_Url string `mapstructure:"CUSTOMER_SVC_URL"`
	Company_Svc_Url  string `mapstructure:"COMPANY_SVC_URL"`
	Product_Svc_Url  string `mapstructure:"PRODUCT_SVC_URL"`
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
