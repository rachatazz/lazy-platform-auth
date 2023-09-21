package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                   string `mapstructure:"GO_AUTH_PORT"`
	JwtTokenSecret         string `mapstructure:"GO_AUTH_JWT_TOKEN_SECRET"`
	JwtRefreshTokenSecret  string `mapstructure:"GO_AUTH_JWT_REFRESH_TOKEN_SECRET"`
	JwtTokenExpired        string `mapstructure:"GO_AUTH_JWT_TOKEN_EXPIRED"`
	JwtRefreshTokenExpired string `mapstructure:"GO_AUTH_JWT_REFRESH_TOKEN_EXPIRED"`
	DataBaseUri            string `mapstructure:"GO_AUTH_DATABASE_URI"`
	ElasticUrl             string `mapstructure:"GO_AUTH_ELASTICE_URL"`
	ElasticApiKey          string `mapstructure:"GO_AUTH_ELASTICE_MAIL_API_KEY"`
}

func ConfigService() (config Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("GO_AUTH_PORT", "3000")
	viper.SetDefault("GO_AUTH_JWT_TOKEN_SECRET", "hello")
	viper.SetDefault("GO_AUTH_JWT_REFRESH_TOKEN_SECRET", "world")
	viper.SetDefault("GO_AUTH_JWT_TOKEN_EXPIRED", "30m")
	viper.SetDefault("GO_AUTH_JWT_REFRESH_TOKEN_EXPIRED", "24h")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}
