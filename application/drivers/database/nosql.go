package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, dbURI string) (*mongo.Client, error) {
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, fmt.Errorf("error while trying to connect to the database: %s", err.Error())
	}

	return cli, nil
}
