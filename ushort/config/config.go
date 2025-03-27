package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Addr string `env:"REDIS_HOST"`
	DB   int    `env:"REDIS_DB"`
}

func LoadConfig(ctx context.Context) *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatalln(err)
	}
	return &cfg
}
