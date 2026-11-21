package usecases

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/infrastructure"
	"fmt"
)

type AnalysisUseCase interface {
	GetFeatures() (*domain.FeaturesResponse, error)
	GetFeatureDistribution(featureName string, bins int) (*domain.FeatureDistribution, error)
	GetMaintenanceGroups() (*domain.MaintenanceGroups, error)
	GetProductionGroups() (*domain.ProductionGroups, error)
	GetRiskGroups() (*domain.RiskGroups, error)
}

type analysisUseCase struct {
	api *infrastructure.APIClient
}

func NewAnalysisUseCase(api *infrastructure.APIClient) AnalysisUseCase {
	return &analysisUseCase{api: api}
}

func (u *analysisUseCase) GetFeatures() (*domain.FeaturesResponse, error) {
	var res domain.FeaturesResponse
	err := u.api.GetJSON("/api/features", &res)
	return &res, err
}

func (u *analysisUseCase) GetFeatureDistribution(featureName string, bins int) (*domain.FeatureDistribution, error) {
	var res domain.FeatureDistribution
	endpoint := fmt.Sprintf("/api/features/%s/distribution?bins=%d", featureName, bins)
	err := u.api.GetJSON(endpoint, &res)
	return &res, err
}

func (u *analysisUseCase) GetMaintenanceGroups() (*domain.MaintenanceGroups, error) {
	var res domain.MaintenanceGroups
	err := u.api.GetJSON("/api/maintenance-groups", &res)
	return &res, err
}

func (u *analysisUseCase) GetProductionGroups() (*domain.ProductionGroups, error) {
	var res domain.ProductionGroups
	err := u.api.GetJSON("/api/production-groups", &res)
	return &res, err
}

func (u *analysisUseCase) GetRiskGroups() (*domain.RiskGroups, error) {
	var res domain.RiskGroups
	err := u.api.GetJSON("/api/risk-groups", &res)
	return &res, err
}
