package config

import "github.com/spf13/viper"

type Config struct {
	Production_Url  string `mapstructure:"PRODUCTION_URL"`
	Development_Url string `mapstructure:"DEVELOPMENT_URL"`
	Port            string `mapstructure:"PORT"`
	DBDriver        string `mapstructure:"DB_DRIVER"`
	DBSource        string `mapstructure:"DB_SOURCE"`
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
