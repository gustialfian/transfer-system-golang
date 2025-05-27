// Package db hold implementation detail on how app will integrate with database
package db

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// MustNewPostgreSQL establishes a connection to a PostgreSQL database using the provided
// user credentials, host, and database name. It applies any pending database migrations
// from the embedded migrations filesystem. If any error occurs during migration or connection,
// the function logs the error and terminates the application.
func MustNewPostgreSQL(user, pass, host, dbname string) *sqlx.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&TimeZone=UTC",
		user, pass, host, dbname,
	)

	s, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", s, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
