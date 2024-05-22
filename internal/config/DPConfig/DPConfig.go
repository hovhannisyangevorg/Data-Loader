package DPConfig

import (
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/DBConfig"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	MinUpdateDate time.Time
	Path          string
	CoreUrl       string
	ApiKey        string
	Database      *DBConfig.DatabaseConfigEnv
}

// Load environment variables from ..env file at the start
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("config: init: ", err)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	return &Config{
		MinUpdateDate: time.Now().AddDate(-1, 0, 0),
		CoreUrl:       getEnv("CORE_URL", ""),
		ApiKey:        getEnv("API_KEY", ""),
		Path:          getEnv("PATH", ""),
		Database: &DBConfig.DatabaseConfigEnv{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnv("DB_PORT", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSLMODE", ""),
		},
	}
}
