package config

import (
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	MinUpdateDate  time.Time
	AnimalsCoreUrl string
}

// TODO: .env sets a path for multiple OSes
func init() {
	// Load environment variables from .env file at the start
	if err := godotenv.Load("C:\\Users\\gev99\\OneDrive\\Desktop\\Data-Procesor\\.env\\"); err != nil {
		log.Println("No .env file found")
	}
}

func LoadConfig() *Config {
	return &Config{
		MinUpdateDate: time.Now(),
	}
}
