package store

import "github.com/redis/go-redis/v9"

type RedisStore struct {
	db *redis.Client
}
