package DataProcesorConfig

import (
	"github.com/hovhannisyangevorg/Data-Procesor/internal/utils"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// Load environment variables from ..env file at the start
func init() {
	if err := godotenv.Load(); err != nil {
		utils.WrapError("config: init: ", err)
	}
}

type ConfigSearch interface {
	GetKey() string
	GetAwsSecret() string
	GetAwsRegion() string
	GetAwsS3BucketName() string
}

type Config struct {
	MinUpdateDate time.Time
	Path          string
	CoreUrl       string
	ApiKey        string
	AWSInternal   AWSConfig
}

type AWSConfig struct {
	AwsKye          string
	AwsSecret       string
	AwsRegion       string
	AwsS3BucketName string
}

func (c *Config) GetAwsKye() string {
	return c.AWSInternal.AwsKye
}

func (c *Config) GetAwsSecret() string {
	return c.AWSInternal.AwsSecret
}

func (c *Config) GetAwsRegion() string {
	return c.AWSInternal.AwsRegion
}

func (c *Config) GetAwsS3BucketName() string {
	return c.AWSInternal.AwsS3BucketName
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
		AWSInternal: AWSConfig{
			AwsKye:          getEnv("AWS_KEY", ""),
			AwsSecret:       getEnv("AWS_SECRET", ""),
			AwsRegion:       getEnv("AWS_REGION", ""),
			AwsS3BucketName: getEnv("AWS_S3_BUCKET_NAME", ""),
		},
	}
}
