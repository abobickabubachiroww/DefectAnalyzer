package domain

type PredictRequest struct {
	ProductionVolume     float64 `json:"ProductionVolume"`
	ProductionCost       float64 `json:"ProductionCost"`
	SupplierQuality      float64 `json:"SupplierQuality"`
	DeliveryDelay        float64 `json:"DeliveryDelay"`
	MaintenanceHours     float64 `json:"MaintenanceHours"`
	DowntimePercentage   float64 `json:"DowntimePercentage"`
	InventoryTurnover    float64 `json:"InventoryTurnover"`
	StockoutRate         float64 `json:"StockoutRate"`
	WorkerProductivity   float64 `json:"WorkerProductivity"`
	SafetyIncidents      int     `json:"SafetyIncidents"`
	EnergyConsumption    float64 `json:"EnergyConsumption"`
	EnergyEfficiency     float64 `json:"EnergyEfficiency"`
	AdditiveProcessTime  float64 `json:"AdditiveProcessTime"`
	AdditiveMaterialCost float64 `json:"AdditiveMaterialCost"`
}

type PredictResponse struct {
	PredictedClass int     `json:"predicted_class"`
	Label          string  `json:"label"`
	Score          float64 `json:"score"`
	Threshold      float64 `json:"threshold"`
	Model          string  `json:"model"`
	Note           string  `json:"note"`
}
