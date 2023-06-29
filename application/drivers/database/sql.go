package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const templateDNS = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

// ConnectSQL uses gorm to provide database connection.
// The DNS template used is 'username:password@tcp(address)/dbname?charset=utf8mb4&parseTime=True&loc=Local'
func ConnectSQL() *gorm.DB {
	DNS := fmt.Sprintf(
		templateDNS,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(DNS))
	if err != nil {
		log.Fatalf("database initialization fail - error: %s", err.Error())
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fail when trying to get master connection - error: %s", err.Error())
	}
	defer sqlDb.Close()

	if err != nil {
		fmt.Println("this")
	}

	return db
}

func DatabasePool(connStr string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v \n", err)
	}
	defer dbpool.Close()

	return dbpool
}
