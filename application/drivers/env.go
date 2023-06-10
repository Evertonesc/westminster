package drivers

import (
	"log"
	"os"
)

const (
	MongoDBPort = "MONGODB_PORT"
	MongoDBHost = "MONGODB_HOST"
)

func EnsureEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("no value for the environment variable %s", key)
	}

	return v
}
