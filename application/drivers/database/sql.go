package database

import "context"

type SqlDatabase interface {
	ConnectSQL(ctx context.Context, databaseURL string) (any, error)
}
