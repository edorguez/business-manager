package config

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	Port        string `mapstructure:"PORT"`
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBSource    string `mapstructure:"DB_SOURCE"`
	Token       string `mapstructure:"TOKEN"`
	Phone_Id    string `mapstructure:"PHONE_ID"`
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
