package handlers

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/usecases"
)

type RecommendationsHandler struct {
	usecase usecases.RecommendationsUseCase
}

func NewRecommendationsHandler(uc usecases.RecommendationsUseCase) *RecommendationsHandler {
	return &RecommendationsHandler{usecase: uc}
}

func (h *RecommendationsHandler) GetGeneralRecommendations() (*domain.GeneralRecommendationsResponse, error) {
	return h.usecase.GetGeneralRecommendations()
}

func (h *RecommendationsHandler) GetPersonalizedRecommendations(req domain.PredictRequest) (*domain.PersonalizedRecommendationsResponse, error) {
	return h.usecase.GetPersonalizedRecommendations(req)
}
