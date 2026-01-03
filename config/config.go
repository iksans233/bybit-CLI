package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey    string
	SecretKey string
	BaseURL   string
}

var C *Config

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Err when load env: %v", err)
	}

	C = &Config{
		ApiKey:    os.Getenv("BYBIT_API_KEY"),
		SecretKey: os.Getenv("BYBIT_SECRET_KEY"),
		BaseURL:   "https://api-demo.bybit.com",
	}
	return nil
}
