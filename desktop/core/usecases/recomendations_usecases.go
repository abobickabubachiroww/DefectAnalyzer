package usecases

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/infrastructure"
)

type RecommendationsUseCase interface {
	GetGeneralRecommendations() (*domain.GeneralRecommendationsResponse, error)
	GetPersonalizedRecommendations(req domain.PredictRequest) (*domain.PersonalizedRecommendationsResponse, error)
}

type recommendationsUseCase struct {
	api *infrastructure.APIClient
}

func NewRecommendationsUseCase(api *infrastructure.APIClient) RecommendationsUseCase {
	return &recommendationsUseCase{api: api}
}

func (u *recommendationsUseCase) GetGeneralRecommendations() (*domain.GeneralRecommendationsResponse, error) {
	var res domain.GeneralRecommendationsResponse
	err := u.api.GetJSON("/api/recommendations", &res)
	return &res, err
}

func (u *recommendationsUseCase) GetPersonalizedRecommendations(req domain.PredictRequest) (*domain.PersonalizedRecommendationsResponse, error) {
	var res domain.PersonalizedRecommendationsResponse
	err := u.api.PostJSON("/api/recommendations/personalized", req, &res)
	return &res, err
}
