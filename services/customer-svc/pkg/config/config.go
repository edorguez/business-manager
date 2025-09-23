package config

import "github.com/spf13/viper"

type Config struct {
	CustomerSvcPort             string `mapstructure:"CUSTOMER_SVC_PORT"`
	PostgresDBDriver            string `mapstructure:"POSTGRES_DB_DRIVER"`
	CustomerDBSourceDevelopment string `mapstructure:"CUSTOMER_DB_SOURCE_DEVELOPMENT"`
	CustomerDBSourceProduction  string `mapstructure:"CUSTOMER_DB_SOURCE_PRODUCTION"`
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
