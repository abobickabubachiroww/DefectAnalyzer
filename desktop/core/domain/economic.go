package domain

type EconomicsRequest struct {
	DowntimeHoursPerMonth    float64 `json:"downtime_hours_per_month"`
	CostPerDowntimeHour      float64 `json:"cost_per_downtime_hour"`
	ExpectedReductionPercent float64 `json:"expected_reduction_percent"`
	ImplementationCost       float64 `json:"implementation_cost"`
	MonthlyOperatingCost     float64 `json:"monthly_operating_cost"`
}

type EconomicsResponse struct {
	GrossMonthlySavings float64 `json:"gross_monthly_savings"`
	NetMonthlySavings   float64 `json:"net_monthly_savings"`
	PaybackMonths       float64 `json:"payback_months"`
	FirstYearNetBenefit float64 `json:"first_year_net_benefit"`
	Currency            string  `json:"currency"`
	Scenario            string  `json:"scenario"`
	Note                string  `json:"note"`
}
