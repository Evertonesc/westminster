package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const templateDNS = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

// ConnectSQL uses gorm to provide database connection.
// The dns template used is 'username:password@tcp(address)/dbname'
func ConnectSQL() *gorm.DB {

	DNS := fmt.Sprintf(templateDNS, os.Getenv("MARIADB_USER"), os.Getenv("MARIADB_PASS"), os.Getenv("MARIADB_HOST"), os.Getenv("MARIADB_DATABASE"))

	db, err := gorm.Open(mysql.Open(DNS))
	if err != nil {
		log.Fatalf("database initialization fail - error: %s", err.Error())
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fail when trying to get master connection - error: %s", err.Error())
	}
	defer sqlDb.Close()

	return db
}
