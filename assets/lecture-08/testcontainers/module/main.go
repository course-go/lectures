package main

import (
	"context"
	"log"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// START OMIT

func main() {
	ctx := context.Background()
	postgresContainer, _ := postgres.RunContainer(ctx, // ignored error
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()
}

// END OMIT
