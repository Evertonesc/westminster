package dbwriters

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MemberWriter struct {
	dbpool *pgxpool.Pool
}

func Members(dbpool *pgxpool.Pool) *MemberWriter {
	return &MemberWriter{
		dbpool: dbpool,
	}
}

// TODO: implement
func (mw *MemberWriter) Create(ctx context.Context) error {
	return nil
}
