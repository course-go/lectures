package main

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// START OMIT

func main() {
	logger := zerolog.New(os.Stderr).With().Str("component", "api").Timestamp().Logger()
	hostname := "localhost:8080"
	logger.Info().
		Str("hostname", hostname).
		Msg("starting server")
	url := "https://course-go.dev"
	logger.Error().
		Str("url", url).
		Dur("timeout", 3*time.Second).
		Err(errors.New("HTTP error")).
		Msg("could not fetch data")
}

// END OMIT
