package usecases

import "DefectAnalyzer/desktop/core/domain"

type StatsUseCase interface {
	GetDashboardStats() (*domain.DashboardStats, error)
	GetModelMetrics() (*domain.ModelMetrics, error)
}

type statsUseCase struct{}

func NewStatsUseCase() StatsUseCase {
	return &statsUseCase{}
}

func (u *statsUseCase) GetDashboardStats() (*domain.DashboardStats, error) {
	return &domain.DashboardStats{
		Total: 3240,
		ClassDistribution: map[string]int{
			"low":  517,
			"high": 2723,
		},
		MaintenanceByClass: map[string]float64{
			"low":  6.79,
			"high": 12.37,
		},
		ProductionVolumeByClass: map[string]float64{
			"low":  470.87,
			"high": 563.27,
		},
		MissingValues: 0,
		Columns:       17,
		ClassShare: domain.ClassShare{
			Low:  0.1596,
			High: 0.8404,
		},
	}, nil
}

func (u *statsUseCase) GetModelMetrics() (*domain.ModelMetrics, error) {
	return &domain.ModelMetrics{
		Model:      "RF Balanced Subsample",
		Threshold:  0.62,
		Evaluation: "previously_evaluated_test_split",
		TestSize:   648,
		Metrics: domain.ModelMetricsData{
			Accuracy:         0.7469,
			PrecisionClass1:  0.9441,
			RecallClass1:     0.7431,
			F1Class1:         0.8316,
			RocAuc:           0.7517,
			BalancedAccuracy: 0.7551,
		},
		ConfusionMatrix: [][]int{
			{79, 24},
			{140, 405},
		},
		Labels: []int{0, 1},
		Note:   "Исторический результат теста; выборка уже использовалась при исследовании моделей",
	}, nil
}
