package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func up(m migrate.Migrate) {
	err := m.Up()
	if err != nil {
		log.Fatalf("failed to uo : %v", err)
	}
}

// to rollback tp specific migration
func rollBack(m migrate.Migrate, n uint) {
	err := m.Migrate(n)
	if err != nil {
		log.Fatalf("failed to uo : %v", err)
	}
}

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:root@localhost:5432/dbname?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	up(*m)

	// rollBack(*m, 2)
	// // Rollback to version 3
	// err = m.Migrate(4)
	// if err != nil {
	// 	log.Fatalf("Failed to migrate to version 3: %v", err)
	// }

	// Rollback the last migration

	// err = m.Steps(-1)
	if err != nil {
		log.Fatalf("Failed to rollback the last migration: %v", err)
	}

	log.Println("Migration reverted successfully")
}
