package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port      string
	DBURL     string
	JWTSecret string
	JWTExpiry int
}

func LoadConfig() (*Config, error) {
	port := getEnv("PORT", "8080")
	dbURL := getEnv("DB_URL", "postgres://postgres:postgres@localhost:5432/terminal?sslmode=disable")
	jwtSecret := getEnv("JWT_SECRET", "very-secret-key")
	jwtExpiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY", "24")) // in hours

	return &Config{
		Port:      port,
		DBURL:     dbURL,
		JWTSecret: jwtSecret,
		JWTExpiry: jwtExpiry,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
