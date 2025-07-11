package config

import (
	"os"
)

type Config struct {
	Port      string
	RedisAddr string
	BaseURL   string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	redisAddr := os.Getenv("REDIS_ADDR")

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	return &Config{
		Port:      port,
		RedisAddr: redisAddr,
		BaseURL:   baseURL,
	}
}
