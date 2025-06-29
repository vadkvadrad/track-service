package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	Db     DbConfig
}

type DbConfig struct {
	Dsn string
}

type AppConfig struct {
	Port string
}

func Load() (*Config, error) {
	err := godotenv.Load(dir(".env"))
	if err != nil {
		log.Println("Error loading .env file, using default config.", "Error:", err.Error())
	}
	config := &Config{
		App: AppConfig{
			Port: getEnv("PORT", "50051"),
		},
		Db: DbConfig{
			Dsn: getEnv("DSN", ""),
		},
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) validate() error {
	if c.Db.Dsn == "" {
		return errors.New("DSN is required")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func dir(envFile string) string {
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Failed to get working directory: %v", err)
    }

    return filepath.Join(currentDir, envFile)
}