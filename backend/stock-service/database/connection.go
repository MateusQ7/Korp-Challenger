package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	DB    *gorm.DB
	rawDB *sql.DB
)

func RawDB() *sql.DB {
	return rawDB
}

func Connect() {
	dsn := "user=postgres password=admin dbname=stock_service sslmode=disable"

	var err error

	rawDB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting (raw) to DB:", err)
	}
	if err = rawDB.Ping(); err != nil {
		log.Fatal("Error pinging DB:", err)
	}

	fmt.Println("Connected (raw) to database.")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting (gorm) to DB:", err)
	}

	fmt.Println("Connected (gorm) to database.")
}
