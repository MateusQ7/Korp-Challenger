package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
