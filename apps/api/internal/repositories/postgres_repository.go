package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct { Pool *pgxpool.Pool }

func NewPostgresRepository(ctx context.Context, url string) (*PostgresRepository, error) { pool, err := pgxpool.New(ctx, url); if err != nil { return nil, err }; return &PostgresRepository{Pool: pool}, nil }
func (r *PostgresRepository) Ping(ctx context.Context) error { return r.Pool.Ping(ctx) }
func (r *PostgresRepository) Close() { r.Pool.Close() }
