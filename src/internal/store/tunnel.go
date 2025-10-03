package store

import (
	"context"
	"errors"
	"log"
	consts "tunnels/tunnels/consts/status"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TunnelStore struct {
	db *pgxpool.Pool
}

func (ts *TunnelStore) GetTunnelFromDomain(ctx context.Context, domainName string) (string, error) {
	var tunnelName string

	sql := `
		SELECT tunnel_name FROM tunnels WHERE domain = $1
	`

	err := ts.db.QueryRow(ctx, sql, domainName).Scan(&tunnelName)
	if err != nil {
		log.Println("Cannot find domain name from the given tunnel name")
		return "", errors.New(consts.CANNOT_GET_DOMAIN)
	}

	return tunnelName, nil
}
