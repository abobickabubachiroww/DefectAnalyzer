package domain

type ClassShare struct {
	Low  float64 `json:"low"`
	High float64 `json:"high"`
}

type DashboardStats struct {
	Total                   int                `json:"total"`
	ClassDistribution       map[string]int     `json:"class_distribution"`
	MaintenanceByClass      map[string]float64 `json:"maintenance_by_class"`
	ProductionVolumeByClass map[string]float64 `json:"production_volume_by_class"`
	MissingValues           int                `json:"missing_values"`
	Columns                 int                `json:"columns"`
	ClassShare              ClassShare         `json:"class_share"`
}

type ModelMetricsData struct {
	Accuracy         float64 `json:"accuracy"`
	PrecisionClass1  float64 `json:"precision_class_1"`
	RecallClass1     float64 `json:"recall_class_1"`
	F1Class1         float64 `json:"f1_class_1"`
	RocAuc           float64 `json:"roc_auc"`
	BalancedAccuracy float64 `json:"balanced_accuracy"`
}

type ModelMetrics struct {
	Model           string           `json:"model"`
	Threshold       float64          `json:"threshold"`
	Evaluation      string           `json:"evaluation"`
	TestSize        int              `json:"test_size"`
	Metrics         ModelMetricsData `json:"metrics"`
	ConfusionMatrix [][]int          `json:"confusion_matrix"`
	Labels          []int            `json:"labels"`
	Note            string           `json:"note"`
}
