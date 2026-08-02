package config

import "testing"

func TestLoadUsesDefaultValues(t *testing.T) {
	for _, key := range []string{"APP_ENV", "PORT", "DATABASE_URL", "REDIS_URL", "CORS_ORIGINS"} {
		t.Setenv(key, "")
	}

	config := Load()
	if config.Environment != "development" || config.Port != "8080" {
		t.Fatalf("unexpected application defaults: %#v", config)
	}
	if config.DatabaseURL == "" || config.RedisURL == "" || config.CorsOrigins == "" {
		t.Fatalf("expected service defaults, got %#v", config)
	}
}

func TestLoadUsesEnvironmentValues(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("PORT", "9000")
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("REDIS_URL", "redis://example")
	t.Setenv("CORS_ORIGINS", "https://example.com")

	config := Load()
	if config != (AppConfig{"production", "9000", "postgres://example", "redis://example", "https://example.com"}) {
		t.Fatalf("unexpected environment configuration: %#v", config)
	}
}
