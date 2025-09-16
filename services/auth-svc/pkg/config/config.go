package config

import "github.com/spf13/viper"

type Config struct {
	Port                    string `mapstructure:"PORT"`
	PostgresDBDriver        string `mapstructure:"POSTGRES_DB_DRIVER"`
	AuthDBSourceDevelopment string `mapstructure:"AUTH_DB_SOURCE_DEVELOPMENT"`
	AuthDBSourceProduction  string `mapstructure:"AUTH_DB_SOURCE_PRODUCTION"`
	JWTSecretKey            string `mapstructure:"JWT_SECRET_KEY"`
	Company_Svc_Url         string `mapstructure:"COMPANY_SVC_URL"`
	Company_Svc_Port        string `mapstructure:"COMPANY_SVC_PORT"`
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
