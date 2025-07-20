package config

import (
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSLMODE,required"`
	JWTSecret  string `env:"JWT_SECRET,required"`
}

func NewEnvConfig() *EnvConfig {
	// Try to load .env file, but don't fail if it doesn't exist (production)
	err := godotenv.Load()
	if err != nil {
		log.Info("No .env file found, using environment variables")
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from environment: %e", err)
	}

	return config
}
