package main

import (
	"errors"
	"time"

	"go.uber.org/zap"
)

// START OMIT

func main() {
	logger, _ := zap.NewProduction() // production preset
	defer logger.Sync()
	logger.Info("failed to fetch data",
		zap.String("URL", "https://course-go.dev"),
		zap.Duration("timeout", 3*time.Second),
		zap.Error(errors.New("HTTP error")),
	)
}

// END OMIT
