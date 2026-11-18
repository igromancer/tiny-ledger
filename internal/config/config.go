package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("Error loading .env file")
	}
}

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
	DbHost      string `env:"DB_HOST"`
	DbUser      string `env:"DB_USER"`
	DbPassword  string `env:"DB_PASSWORD"`
}

func Get() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Println("Unable to parse the environment, falling back to defaults.")
	}
	log.Println(cfg)
	return &cfg
}
