package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct { Client *redis.Client }

func NewRedisRepository(url string) (*RedisRepository, error) { options, err := redis.ParseURL(url); if err != nil { return nil, err }; return &RedisRepository{Client: redis.NewClient(options)}, nil }
func (r *RedisRepository) Ping(ctx context.Context) error { return r.Client.Ping(ctx).Err() }
func (r *RedisRepository) Close() error { return r.Client.Close() }
