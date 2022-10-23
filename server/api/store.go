package api

import (
	"os"

	"github.com/alekstet/linux_service_checker/server/db"
	"github.com/alekstet/linux_service_checker/server/maker"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rs/zerolog"
)

type store struct {
	Log   zerolog.Logger
	maker maker.Maker
}

type Getter interface {
	getData() (*maker.ServicesInfo, error)
}

type PGXPool struct {
	DBPool *pgxpool.Pool
}

func NewPGXPool() (*PGXPool, error) {
	pool, err := db.GetDBConnectionPool()
	if err != nil {
		return nil, err
	}

	return &PGXPool{
		DBPool: pool,
	}, nil
}

func NewStore(maker maker.Maker) *store {
	return &store{
		Log:   zerolog.New(os.Stdout).With().Timestamp().Logger(),
		maker: maker,
	}
}
