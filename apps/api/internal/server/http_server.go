package server

import (
	"context"
	"fmt"
	"github.com/car-rental-template/api/internal/config"
	"github.com/car-rental-template/api/internal/controllers"
	"github.com/car-rental-template/api/internal/middleware"
	"github.com/car-rental-template/api/internal/repositories"
	"github.com/car-rental-template/api/internal/routes"
	"github.com/car-rental-template/api/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct{ http *http.Server }

func New(cfg config.AppConfig) (*Server, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	postgres, err := repositories.NewPostgresRepository(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, nil, err
	}
	redis, err := repositories.NewRedisRepository(cfg.RedisURL)
	if err != nil {
		postgres.Close()
		return nil, nil, err
	}
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), cors.New(middleware.Cors(cfg.CorsOrigins)))
	health := controllers.NewHealthController(services.NewHealthService(postgres, redis))
	routes.Register(router, health)
	return &Server{http: &http.Server{Addr: ":" + cfg.Port, Handler: router, ReadHeaderTimeout: 10 * time.Second}}, func() { postgres.Close(); _ = redis.Close() }, nil
}
func (s *Server) Start() {
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Errorf("server failed: %w", err))
	}
}
func (s *Server) Stop(ctx context.Context) error { return s.http.Shutdown(ctx) }
