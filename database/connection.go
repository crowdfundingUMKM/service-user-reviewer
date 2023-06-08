package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnectionDB() *gorm.DB {
	// initial connected db
	var (
		dbUser    = os.Getenv("DB_USER")       // e.g. 'my-db-user'
		dbPwd     = os.Getenv("DB_PASS")       // e.g. 'my-db-password'
		dbName    = os.Getenv("DB_NAME")       // e.g. 'my-database'
		dbPort    = os.Getenv("DB_PORT")       // e.g. '3306'
		dbTCPHost = os.Getenv("INSTANCE_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	// end connected

	// open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	return db

}
