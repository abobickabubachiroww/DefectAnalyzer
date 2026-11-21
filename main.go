package main

import (
	"DefectAnalyzer/desktop/handlers"
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var assets embed.FS

func main() {

	appHandler := &handlers.AppHandler{}

	app := application.New(application.Options{
		Name:        "DefectAnalyzer",
		Description: "Система анализа брака",
		Services: []application.Service{
			application.NewService(appHandler),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:      "DefectAnalyzer",
		Width:      1600,
		Height:     768,
		StartState: application.WindowStateMaximised,
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
