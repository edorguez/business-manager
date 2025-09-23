package config

import "github.com/spf13/viper"

type Config struct {
	OrderSvcPort    string `mapstructure:"ORDER_SVC_PORT"`
	CustomerSvcUrl  string `mapstructure:"CUSTOMER_SVC_URL"`
	CustomerSvcPort string `mapstructure:"CUSTOMER_SVC_PORT"`
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
