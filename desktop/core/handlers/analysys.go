package handlers

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/usecases"
)

type AnalysisHandler struct {
	usecase usecases.AnalysisUseCase
}

func NewAnalysisHandler(uc usecases.AnalysisUseCase) *AnalysisHandler {
	return &AnalysisHandler{usecase: uc}
}

func (h *AnalysisHandler) GetFeatures() (*domain.FeaturesResponse, error) {
	return h.usecase.GetFeatures()
}

func (h *AnalysisHandler) GetFeatureDistribution(featureName string, bins int) (*domain.FeatureDistribution, error) {
	return h.usecase.GetFeatureDistribution(featureName, bins)
}

func (h *AnalysisHandler) GetMaintenanceGroups() (*domain.MaintenanceGroups, error) {
	return h.usecase.GetMaintenanceGroups()
}

func (h *AnalysisHandler) GetProductionGroups() (*domain.ProductionGroups, error) {
	return h.usecase.GetProductionGroups()
}

func (h *AnalysisHandler) GetRiskGroups() (*domain.RiskGroups, error) {
	return h.usecase.GetRiskGroups()
}
