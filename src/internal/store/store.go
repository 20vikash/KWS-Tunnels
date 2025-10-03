package store

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	TunnelStore interface {
	}

	InMemoryStore interface {
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
