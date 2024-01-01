package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_URL      string
	DB_DATABASE string
	PORT        string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}

	log.Println("🚀 Config successfully connected")
}
