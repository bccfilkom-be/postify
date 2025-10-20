package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string `env:"APP_ENV,required"`
	AppPort string `env:"APP_PORT,required"`
}

var cfg Config

func Load() Config {
	if cfg.AppEnv == "" {
		return newConfig()
	}

	return cfg
}

func newConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		if os.IsNotExist(err) {
			// .env file does not exist, proceed with system env variables
			os.Getenv("SHELL")
		} else {
			log.Fatal("Error loading .env file")
		}
	}

	// parse
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}
