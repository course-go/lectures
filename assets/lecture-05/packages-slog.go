package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

// START OMIT

func main() {
	textLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	textLogger.Info("Info message")
	textLogger.Error("Error message")
	fmt.Println()

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(jsonLogger)
	slog.Debug("Debug message") // Debug is ignored by default
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
	fmt.Println()

	slog.Info(
		"HTTP request",
		"method", "POST",
		"status", 201,
	)

	err := errors.New("ran out of gophers")
	slog.ErrorContext(context.Background(), "upload failed", slog.Any("error", err))
}

// END OMIT
