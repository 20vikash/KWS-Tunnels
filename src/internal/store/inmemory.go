package store

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	db *redis.Client
}

func (rs *RedisStore) GetUidFromSecret(ctx context.Context, secret string) (int, error) {
	val, err := rs.db.Get(ctx, secret).Int()
	if err != nil {
		log.Println("Failed to get tunnel UID int")
		return -1, err
	}

	return val, nil
}
