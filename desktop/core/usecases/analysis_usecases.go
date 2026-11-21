package usecases

import "DefectAnalyzer/desktop/core/domain"

type AnalysisUseCase interface {
	GetFeatures() (*domain.FeaturesResponse, error)
	GetFeatureDistribution(featureName string, bins int) (*domain.FeatureDistribution, error)
	GetMaintenanceGroups() (*domain.MaintenanceGroups, error)
	GetProductionGroups() (*domain.ProductionGroups, error)
	GetRiskGroups() (*domain.RiskGroups, error)
}

type analysisUseCase struct{}

func NewAnalysisUseCase() AnalysisUseCase {
	return &analysisUseCase{}
}

func (u *analysisUseCase) GetFeatures() (*domain.FeaturesResponse, error) {
	return &domain.FeaturesResponse{
		Features: []domain.Feature{
			{Name: "ProductionVolume", Label: "Объём производства", Unit: "ед./день"},
			{Name: "MaintenanceHours", Label: "Обслуживание оборудования", Unit: "ч/неделю"},
			{Name: "SupplierQuality", Label: "Качество поставщиков", Unit: "%"},
		},
		Target: "DefectStatus",
	}, nil
}

func (u *analysisUseCase) GetFeatureDistribution(featureName string, bins int) (*domain.FeatureDistribution, error) {
	// В будущем здесь будет запрос к Python с передачей featureName и bins
	return &domain.FeatureDistribution{
		Feature:  featureName,
		Label:    "Обслуживание оборудования",
		Unit:     "ч/неделю",
		BinEdges: []float64{0, 5, 10, 15},
		Series: []domain.ClassCount{
			{Class: 0, Counts: []int{200, 250, 67}},
			{Class: 1, Counts: []int{100, 800, 1823}},
		},
		Note: "Количество наблюдений по интервалам; класс 0 — низкая, 1 — высокая дефектность",
	}, nil
}

func (u *analysisUseCase) GetMaintenanceGroups() (*domain.MaintenanceGroups, error) {
	return &domain.MaintenanceGroups{
		ThresholdHoursPerWeek: 11,
		BelowThreshold: domain.GroupStats{
			Total:           1200,
			HighDefect:      840,
			LowDefect:       360,
			HighDefectShare: 0.7,
		},
		AtOrAboveThreshold: domain.GroupStats{
			Total:           2040,
			HighDefect:      1883,
			LowDefect:       157,
			HighDefectShare: 0.923,
		},
	}, nil
}

func (u *analysisUseCase) GetProductionGroups() (*domain.ProductionGroups, error) {
	return &domain.ProductionGroups{
		ThresholdUnitsPerDay: 800,
		BelowThreshold: domain.GroupStats{
			Total:           2900,
			HighDefect:      2400,
			LowDefect:       500,
			HighDefectShare: 0.8276,
		},
		AtOrAboveThreshold: domain.GroupStats{
			Total:           340,
			HighDefect:      323,
			LowDefect:       17,
			HighDefectShare: 0.95,
		},
	}, nil
}

func (u *analysisUseCase) GetRiskGroups() (*domain.RiskGroups, error) {
	return &domain.RiskGroups{
		Thresholds: domain.RiskThresholds{
			MaintenanceHoursPerWeek: 11,
			ProductionUnitsPerDay:   800,
		},
		Groups: []domain.RiskGroupItem{
			{MaintenanceGe11: false, VolumeGe800: false, Total: 900, HighDefect: 600, LowDefect: 300, HighDefectShare: 0.6667},
			{MaintenanceGe11: false, VolumeGe800: true, Total: 100, HighDefect: 90, LowDefect: 10, HighDefectShare: 0.9},
			{MaintenanceGe11: true, VolumeGe800: false, Total: 2000, HighDefect: 1800, LowDefect: 200, HighDefectShare: 0.9},
			{MaintenanceGe11: true, VolumeGe800: true, Total: 240, HighDefect: 223, LowDefect: 17, HighDefectShare: 0.929},
		},
		Note: "Доля класса 1 среди наблюдений группы, а не доля бракованных изделий",
	}, nil
}
