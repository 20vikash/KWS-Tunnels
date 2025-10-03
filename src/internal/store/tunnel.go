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

func (ts *TunnelStore) ValidateTunnelFromUID(ctx context.Context, uid int, tunnel string) (bool, error) {
	var tunnelName string

	sql := `
		SELECT tunnel_name FROM tunnels WHERE user_id = $1
	`

	rows, err := ts.db.Query(ctx, sql, uid)
	if err != nil {
		log.Println("Cannot get tunnels from the given UID")
		return false, errors.New(consts.CANNOT_GET_TUNNEL)
	}

	for rows.Next() {
		err = rows.Scan(&tunnelName)
		if err != nil {
			log.Println("Error scanning the tunnels tables")
			return false, errors.New(consts.CANNOT_GET_TUNNEL)
		}
		if tunnelName == tunnel {
			return true, nil
		}
	}

	return false, errors.New(consts.NO_TUNNEL_AGAINST_UID)
}

func (ts *TunnelStore) GetDomainFromTunnel(ctx context.Context, tunnelName string, uid int) (string, error) {
	var domainName string

	sql := `
		SELECT domain FROM tunnels WHERE tunnel_name = $1 AND user_id = $2
	`

	err := ts.db.QueryRow(ctx, sql, tunnelName, uid).Scan(&domainName)
	if err != nil {
		log.Println("Cannot find domain name from the given tunnel name")
		return "", errors.New(consts.CANNOT_GET_DOMAIN)
	}

	return tunnelName, nil
}
