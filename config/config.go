package config

import (
	"log"
	"os"
)

var (
	Port      string
	AuthToken string
)

func Load() {
	Port = getEnv("PORT", "8080")
	AuthToken = getEnv("AUTH_TOKEN", "abcd")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func init() {
	Load()
	log.Printf("Configuration loaded: Port=%s, AuthToken=%s", Port, AuthToken)
}
