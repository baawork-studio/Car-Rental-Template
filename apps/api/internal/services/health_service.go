package services

import (
	"context"
	"github.com/car-rental-template/api/internal/models"
	"github.com/car-rental-template/api/internal/repositories"
)

type HealthService struct { postgres *repositories.PostgresRepository; redis *repositories.RedisRepository }
func NewHealthService(postgres *repositories.PostgresRepository, redis *repositories.RedisRepository) *HealthService { return &HealthService{postgres: postgres, redis: redis} }
func (s *HealthService) Status(ctx context.Context) models.HealthStatus { database, cache := "ok", "ok"; if s.postgres.Ping(ctx) != nil { database = "unavailable" }; if s.redis.Ping(ctx) != nil { cache = "unavailable" }; status := "ok"; if database != "ok" || cache != "ok" { status = "degraded" }; return models.HealthStatus{Status: status, Database: database, Cache: cache} }
