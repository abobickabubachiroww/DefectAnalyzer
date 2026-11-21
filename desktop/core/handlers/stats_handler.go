package handlers

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/usecases"
)

type StatsHandler struct {
	usecase usecases.StatsUseCase
}

func NewStatsHandler(uc usecases.StatsUseCase) *StatsHandler {
	return &StatsHandler{usecase: uc}
}

func (h *StatsHandler) GetDashboardStats() (*domain.DashboardStats, error) {
	return h.usecase.GetDashboardStats()
}

func (h *StatsHandler) GetModelMetrics() (*domain.ModelMetrics, error) {
	return h.usecase.GetModelMetrics()
}
