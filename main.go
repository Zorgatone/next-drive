package main

import (
	"context"
	"embed"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	appLogger := logger.NewDefaultLogger()

	// Create an instance of the app structure
	app := NewApp(appLogger)

	var wailsApp *application.Application = nil

	wailsApp = application.NewWithOptions(&options.App{
		Title:  "next-drive",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				FullSizeContent: true,
			},
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			err := app.startup(ctx)

			if err != nil {
				appLogger.Error(fmt.Sprintf("startup error: %s", err.Error()))
				wailsApp.Quit()
			}
		},
		Bind: []interface{}{
			app,
		},
	})

	// Create application with options
	err := wailsApp.Run()

	if err != nil {
		appLogger.Error(fmt.Sprintf("run error: %s", err.Error()))
	}
}
