package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBConnectionPool(connectionString string) (*pgxpool.Pool, error) {
	connectionPool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return connectionPool, nil
}
