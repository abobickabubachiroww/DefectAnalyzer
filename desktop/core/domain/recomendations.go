package domain

type GeneralRecommendation struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Basis   string   `json:"basis"`
	Actions []string `json:"actions"`
}

type GeneralRecommendationsResponse struct {
	Recommendations []GeneralRecommendation `json:"recommendations"`
	Disclaimer      string                  `json:"disclaimer"`
}

type PersonalizedRecommendation struct {
	Feature     string   `json:"feature"`
	Value       float64  `json:"value"`
	ShapValue   float64  `json:"shap_value"`
	Priority    string   `json:"priority"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Actions     []string `json:"actions"`
}

type PersonalizedRecommendationsResponse struct {
	PredictedClass  int                          `json:"predicted_class"`
	Label           string                       `json:"label"`
	Score           float64                      `json:"score"`
	Recommendations []PersonalizedRecommendation `json:"recommendations"`
	Disclaimer      string                       `json:"disclaimer"`
}
