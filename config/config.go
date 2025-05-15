package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL      string
	UsersServiceHost string
	UsersServicePort int
	PostsServiceHost string
	PostsServicePort int
	ChatServiceHost  string
	ChatServicePort  int
	LogLevel         string
}

// LoadConfig loads the configuration from environment variables and .env file.
func LoadConfig() (*Config, error) {
	err := godotenv.Load() // Load .env file if it exists
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	config := &Config{
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://user:password@host:port/database?sslmode=disable"),
		UsersServiceHost: getEnv("USERS_SERVICE_HOST", "localhost"),
		PostsServiceHost: getEnv("POSTS_SERVICE_HOST", "localhost"),
		ChatServiceHost:  getEnv("CHAT_SERVICE_HOST", "localhost"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
	}

	config.UsersServicePort, err = getEnvAsInt("USERS_SERVICE_PORT", 8080)
	if err != nil {
		return nil, err
	}

	config.PostsServicePort, err = getEnvAsInt("POSTS_SERVICE_PORT", 8081)
	if err != nil {
		return nil, err
	}

	config.ChatServicePort, err = getEnvAsInt("CHAT_SERVICE_PORT", 8082)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) (int, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue, nil
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err // Or handle the error as appropriate
	}
	return value, nil
}
