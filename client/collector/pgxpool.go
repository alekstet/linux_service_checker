package collector

import (
	"github.com/alekstet/linux_service_checker/client/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ Collector = (*PGXPool)(nil)

type PGXPool struct {
	DBPool *pgxpool.Pool
}

func NewPGXPool(connectionString string) (*PGXPool, error) {
	pool, err := db.GetDBConnectionPool(connectionString)
	if err != nil {
		return nil, err
	}

	return &PGXPool{
		DBPool: pool,
	}, nil
}
