package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDBConnectionPool() (*pgxpool.Pool, error) {
	connectionString := "postgresql://postgres:21narufu@localhost:5432/postgres?sslmode=disable"
	connectionPool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	return connectionPool, nil
}
