package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	_ "github.com/mattn/go-sqlite3"
)

// START OMIT

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	defer db.Close()
	checker := health.NewChecker(
		health.WithCheck(health.Check{
			Name:    "database",
			Timeout: 2 * time.Second,
			Check:   db.PingContext,
		}),
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name: "search",
			Check: func(ctx context.Context) error {
				return fmt.Errorf("this makes the check fail")
			},
		}),
	)
	http.Handle("/health", health.NewHandler(checker))
	http.ListenAndServe(":3000", nil)
}

// END OMIT
