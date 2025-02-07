package config

import "github.com/spf13/viper"

type Config struct {
	Environment         string `mapstructure:"ENVIRONMENT"`
	Port                string `mapstructure:"PORT"`
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	Twilio_Account_SID  string `mapstructure:"TWILIO_ACCOUNT_SID"`
	Twilio_Auth_Token   string `mapstructure:"TWILIO_AUTH_TOKEN"`
	Twilio_Phone_Number string `mapstructure:"TWILIO_PHONE_NUMBER"`
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
