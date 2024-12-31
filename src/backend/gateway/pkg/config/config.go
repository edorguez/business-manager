package config

import "github.com/spf13/viper"

type Config struct {
	Production_Url    string `mapstructure:"PRODUCTION_URL"`
	Development_Url   string `mapstructure:"DEVELOPMENT_URL"`
	Gateway_Port      string `mapstructure:"GATEWAY_PORT"`
	Customer_Svc_Url  string `mapstructure:"CUSTOMER_SVC_URL"`
	Customer_Svc_Port string `mapstructure:"CUSTOMER_SVC_PORT"`
	Company_Svc_Url   string `mapstructure:"COMPANY_SVC_URL"`
	Company_Svc_Port  string `mapstructure:"COMPANY_SVC_PORT"`
	Product_Svc_Url   string `mapstructure:"PRODUCT_SVC_URL"`
	Product_Svc_Port  string `mapstructure:"PRODUCT_SVC_PORT"`
	Auth_Svc_Url      string `mapstructure:"AUTH_SVC_URL"`
	Auth_Svc_Port     string `mapstructure:"AUTH_SVC_PORT"`
	Order_Svc_Url     string `mapstructure:"ORDER_SVC_URL"`
	Order_Svc_Port    string `mapstructure:"ORDER_SVC_PORT"`
	File_Svc_Url      string `mapstructure:"FILE_SVC_URL"`
	File_Svc_Port     string `mapstructure:"FILE_SVC_PORT"`
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
