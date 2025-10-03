package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	TunnelStore interface {
		ValidateTunnelFromUID(ctx context.Context, uid int, tunnel string) (bool, error)
		GetDomainFromTunnel(ctx context.Context, tunnelName string, uid int) (string, error)
	}

	InMemoryStore interface {
		GetUidFromSecret(ctx context.Context, secret string) (int, error)
	}
}

func NewStore(pg *pgxpool.Pool, redis *redis.Client) *Storage {
	return &Storage{
		TunnelStore: &TunnelStore{
			db: pg,
		},
		InMemoryStore: &RedisStore{
			db: redis,
		},
	}
}
