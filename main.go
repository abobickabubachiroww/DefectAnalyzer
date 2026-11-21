package main

import (
	"DefectAnalyzer/desktop/core/handlers"
	"DefectAnalyzer/desktop/core/infrastructure" // 👈 Импортируем
	"DefectAnalyzer/desktop/core/usecases"
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Инициализируем API-клиент (указываем порт Python)
	apiClient := infrastructure.NewAPIClient("http://127.0.0.1:8001")

	// Передаем клиент в каждый UseCase
	statsUseCase := usecases.NewStatsUseCase(apiClient)
	statsHandler := handlers.NewStatsHandler(statsUseCase)

	analysisUseCase := usecases.NewAnalysisUseCase(apiClient)
	analysisHandler := handlers.NewAnalysisHandler(analysisUseCase)

	predictUseCase := usecases.NewPredictUseCase(apiClient)
	predictHandler := handlers.NewPredictHandler(predictUseCase)

	recUseCase := usecases.NewRecommendationsUseCase(apiClient)
	recHandler := handlers.NewRecommendationsHandler(recUseCase)

	economicsUseCase := usecases.NewEconomicsUseCase(apiClient)
	economicsHandler := handlers.NewEconomicsHandler(economicsUseCase)

	appHandler := &handlers.AppHandler{}

	app := application.New(application.Options{
		Name:        "DefectAnalyzer",
		Description: "Система анализа брака",
		Services: []application.Service{
			application.NewService(appHandler),
			application.NewService(statsHandler),
			application.NewService(analysisHandler),
			application.NewService(predictHandler),
			application.NewService(recHandler),
			application.NewService(economicsHandler),
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
