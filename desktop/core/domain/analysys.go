package domain

type Feature struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Unit  string `json:"unit"`
}

type FeaturesResponse struct {
	Features []Feature `json:"features"`
	Target   string    `json:"target"`
}

type ClassCount struct {
	Class  int   `json:"class"`
	Counts []int `json:"counts"`
}

type FeatureDistribution struct {
	Feature  string       `json:"feature"`
	Label    string       `json:"label"`
	Unit     string       `json:"unit"`
	BinEdges []float64    `json:"bin_edges"`
	Series   []ClassCount `json:"series"`
	Note     string       `json:"note"`
}

// Структура для групп
type GroupStats struct {
	Total           int     `json:"total"`
	HighDefect      int     `json:"high_defect"`
	LowDefect       int     `json:"low_defect"`
	HighDefectShare float64 `json:"high_defect_share"`
}

// Группы по обслуживанию
type MaintenanceGroups struct {
	ThresholdHoursPerWeek int        `json:"threshold_hours_per_week"`
	BelowThreshold        GroupStats `json:"below_threshold"`       // Теперь GroupStats
	AtOrAboveThreshold    GroupStats `json:"at_or_above_threshold"` // Теперь GroupStats
}

// Группы по объему производства
type ProductionGroups struct {
	ThresholdUnitsPerDay int        `json:"threshold_units_per_day"`
	BelowThreshold       GroupStats `json:"below_threshold"`       // Теперь GroupStats
	AtOrAboveThreshold   GroupStats `json:"at_or_above_threshold"` // Теперь GroupStats
}

// Структуры для совместного риска
type RiskThresholds struct {
	MaintenanceHoursPerWeek int `json:"maintenance_hours_per_week"`
	ProductionUnitsPerDay   int `json:"production_units_per_day"`
}

type RiskGroupItem struct {
	MaintenanceGe11 bool    `json:"maintenance_ge_11"`
	VolumeGe800     bool    `json:"volume_ge_800"`
	Total           int     `json:"total"`
	HighDefect      int     `json:"high_defect"`
	LowDefect       int     `json:"low_defect"`
	HighDefectShare float64 `json:"high_defect_share"`
}

type RiskGroups struct {
	Thresholds RiskThresholds  `json:"thresholds"`
	Groups     []RiskGroupItem `json:"groups"`
	Note       string          `json:"note"`
}
