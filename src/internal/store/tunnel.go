package store

import "github.com/jackc/pgx/v5/pgxpool"

type TunnelStore struct {
	db *pgxpool.Pool
}
