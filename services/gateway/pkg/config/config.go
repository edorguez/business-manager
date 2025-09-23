package config

import "github.com/spf13/viper"

type Config struct {
	ProductionUrl   string `mapstructure:"PRODUCTION_URL"`
	DevelopmentUrl  string `mapstructure:"DEVELOPMENT_URL"`
	GatewayPort     string `mapstructure:"GATEWAY_PORT"`
	CustomerSvcUrl  string `mapstructure:"CUSTOMER_SVC_URL"`
	CustomerSvcPort string `mapstructure:"CUSTOMER_SVC_PORT"`
	CompanySvcUrl   string `mapstructure:"COMPANY_SVC_URL"`
	CompanySvcPort  string `mapstructure:"COMPANY_SVC_PORT"`
	ProductSvcUrl   string `mapstructure:"PRODUCT_SVC_URL"`
	ProductSvcPort  string `mapstructure:"PRODUCT_SVC_PORT"`
	AuthSvcUrl      string `mapstructure:"AUTH_SVC_URL"`
	AuthSvcPort     string `mapstructure:"AUTH_SVC_PORT"`
	OrderSvcUrl     string `mapstructure:"ORDER_SVC_URL"`
	OrderSvcPort    string `mapstructure:"ORDER_SVC_PORT"`
	FileSvcUrl      string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort     string `mapstructure:"FILE_SVC_PORT"`
	WhatsappSvcUrl  string `mapstructure:"WHATSAPP_SVC_URL"`
	WhatsappSvcPort string `mapstructure:"WHATSAPP_SVC_PORT"`
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
