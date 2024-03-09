package sqldb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DatabasePool(ctx context.Context, connStr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("parsing postgres connection string: %s/n", err.Error())
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("creating connection pool: %s", err.Error())
	}

	defer dbpool.Close()

	return dbpool
}
