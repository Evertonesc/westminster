package drivers

import (
	"log"
	"os"
)

const (
	PostgresConnString = "POSTGRES_CONN_STRING"
)

func EnsureEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("no value for the environment variable %s", key)
	}

	return v
}
