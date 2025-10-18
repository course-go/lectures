package main

import (
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// START OMIT

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := &App{}
	wails.Run(&options.App{
		Title:       "My awesome app",
		AssetServer: &assetserver.Options{Assets: assets},
		OnStartup:   app.startup,
		Bind:        []any{app},
	})
}

type App struct {
	ctx context.Context
}

func (a *App) Sum(x, y int) int {
	return x + y
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// END OMIT
