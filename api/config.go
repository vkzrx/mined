package api

import (
	"strings"

	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	Cors cors.Options
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// TODO handle error
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		}
	}

	return &Config{
		Port: viper.GetString("PORT"),
		Cors: LoadCorsOptions(),
	}
}

func LoadCorsOptions() cors.Options {
	allowedOrigins := viper.GetString("ALLOWED_ORIGINS")
	return cors.Options{
		AllowedOrigins: strings.Split(allowedOrigins, ","),
	}
}
