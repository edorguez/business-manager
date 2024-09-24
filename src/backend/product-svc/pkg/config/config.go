package config

import "github.com/spf13/viper"

type Config struct {
	Environment     string `mapstructure:"ENVIRONMENT"`
	Port            string `mapstructure:"PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBSource        string `mapstructure:"DB_SOURCE"`
	Company_Svc_Url string `mapstructure:"COMPANY_SVC_URL"`
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
