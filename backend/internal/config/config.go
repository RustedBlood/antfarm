package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Database_url string
	Migrate_dir  string
}

func LoadConfig() (*DbConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &DbConfig{
		Database_url: getEnvStr("DATABASE_URL"),
		Migrate_dir:  getEnvStr("MIGRATE_DIR"),
	}, nil
}

func getEnvStr(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic(fmt.Sprintf("Failed to get value with key %s. Check .env", key))
}
