package api

import (
	"log"
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
		log.Fatalf("Error while reading config file %s", err)
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
