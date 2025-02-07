package config

import (
	"log"
	"os"
)

var (
	Port          string
	AuthToken     string
	RedisUrl      string
	RedisPassword string
	RedisQueue    string
	PluginsPath   string
)

func Load() {
	Port = getEnv("PORT", "8089")
	AuthToken = getEnv("AUTH_TOKEN", "abcd")
	RedisUrl = getEnv("REDIS_URL", "localhost:6379")
	RedisPassword = getEnv("REDIS_PASSWORD", "")
	RedisQueue = getEnv("REDIS_QUEUE", "store")
	PluginsPath = getEnv("PLUGINS_PATH", "tasks.yaml")
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
