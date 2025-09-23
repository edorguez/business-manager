package config

import "github.com/spf13/viper"

type Config struct {
	WhatsappSvcPort             string `mapstructure:"WHATSAPP_SVC_PORT"`
	PostgresDBDriver            string `mapstructure:"POSTGRES_DB_DRIVER"`
	WhatsappDBSourceDevelopment string `mapstructure:"WHATSAPP_DB_SOURCE_DEVELOPMENT"`
	WhatsappDBSourceProduction  string `mapstructure:"WHATSAPP_DB_SOURCE_PRODUCTION"`
	TwilioAccountSID            string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken             string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioPhoneNumber           string `mapstructure:"TWILIO_PHONE_NUMBER"`
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
