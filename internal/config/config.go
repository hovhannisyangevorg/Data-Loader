package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	MinUpdateDate time.Time
	CoreUrl       string
	Path          string
	Database      DatabaseConfigEnv
}

// Load environment variables from ..env file at the start
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No ..env file found")
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		if fallback == "" {
			log.Fatalf("Environment variable %s is not set", key)
		}
		return fallback
	}
	return value
}

func LoadConfig() *Config {
	return &Config{
		MinUpdateDate: time.Now().AddDate(-1, 0, 0),
		CoreUrl:       getEnv("CORE_URL", ""),
		Path:          getEnv("PATH", ""),
		Database: DatabaseConfigEnv{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSLMODE", ""),
		},
	}
}
