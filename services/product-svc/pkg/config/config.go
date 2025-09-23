package config

import "github.com/spf13/viper"

type Config struct {
	ProductSvcPort             string `mapstructure:"PRODUCT_SVC_PORT"`
	ProductDBName              string `mapstructure:"PRODUCT_DB_NAME"`
	ProductDBSourceDevelopment string `mapstructure:"PRODUCT_DB_SOURCE_DEVELOPMENT"`
	ProductDBSourceProduction  string `mapstructure:"PRODUCT_DB_SOURCE_PRODUCTION"`
	CompanySvcUrl              string `mapstructure:"COMPANY_SVC_URL"`
	CompanySvcPort             string `mapstructure:"COMPANY_SVC_PORT"`
	FileSvcUrl                 string `mapstructure:"FILE_SVC_URL"`
	FileSvcPort                string `mapstructure:"FILE_SVC_PORT"`
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
