package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=admin dbname=billing_service sslmode=disable")
	if err != nil {
		log.Fatal("Error while connecting to the database:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Connected to the database successfully!")
	return DB
}
