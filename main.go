package main

import (
	"DefectAnalyzer/desktop/core/handlers"
	"DefectAnalyzer/desktop/core/usecases"
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Инициализация Stats (Главная)
	statsUseCase := usecases.NewStatsUseCase()
	statsHandler := handlers.NewStatsHandler(statsUseCase)

	// Инициализация Analysis (Анализ)
	analysisUseCase := usecases.NewAnalysisUseCase()
	analysisHandler := handlers.NewAnalysisHandler(analysisUseCase)

	appHandler := &handlers.AppHandler{}

	app := application.New(application.Options{
		Name:        "DefectAnalyzer",
		Description: "Система анализа брака",
		Services: []application.Service{
			application.NewService(appHandler),
			application.NewService(statsHandler),    // Регистрируем первый
			application.NewService(analysisHandler), // Регистрируем второй
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

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
