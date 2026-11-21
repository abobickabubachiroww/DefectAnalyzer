package usecases

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/infrastructure"
)

type StatsUseCase interface {
	GetDashboardStats() (*domain.DashboardStats, error)
	GetModelMetrics() (*domain.ModelMetrics, error)
}

type statsUseCase struct {
	api *infrastructure.APIClient
}

func NewStatsUseCase(api *infrastructure.APIClient) StatsUseCase {
	return &statsUseCase{api: api}
}

func (u *statsUseCase) GetDashboardStats() (*domain.DashboardStats, error) {
	var res domain.DashboardStats
	err := u.api.GetJSON("/api/dashboard", &res)
	return &res, err
}

func (u *statsUseCase) GetModelMetrics() (*domain.ModelMetrics, error) {
	var res domain.ModelMetrics
	err := u.api.GetJSON("/api/model-metrics", &res)
	return &res, err
}
