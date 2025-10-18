package main

// START OMIT

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
)

func main() {
	m, err := migrate.New(
		"files:///migrations",
		"postgres://localhost:5432/database?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
