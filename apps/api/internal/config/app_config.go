package config

import "os"

type AppConfig struct { Environment, Port, DatabaseURL, RedisURL, CorsOrigins string }

func Load() AppConfig {
	return AppConfig{Environment: value("APP_ENV", "development"), Port: value("PORT", "8080"), DatabaseURL: value("DATABASE_URL", "postgres://car_rental:car_rental@localhost:5432/car_rental?sslmode=disable"), RedisURL: value("REDIS_URL", "redis://localhost:6379/0"), CorsOrigins: value("CORS_ORIGINS", "http://localhost:5173,http://localhost:5174")}
}

func value(key, fallback string) string { if v := os.Getenv(key); v != "" { return v }; return fallback }
