package config

import "github.com/spf13/viper"

type Config struct {
	Port                       string `mapstructure:"PORT"`
	PostgresDBDriver           string `mapstructure:"POSTGRES_DB_DRIVER"`
	CompanyDBSourceDevelopment string `mapstructure:"COMPANY_DB_SOURCE_DEVELOPMENT"`
	CompanyDBSourceProduction  string `mapstructure:"COMPANY_DB_SOURCE_PRODUCTION"`
	File_Svc_Url               string `mapstructure:"FILE_SVC_URL"`
	File_Svc_Port              string `mapstructure:"FILE_SVC_PORT"`
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
