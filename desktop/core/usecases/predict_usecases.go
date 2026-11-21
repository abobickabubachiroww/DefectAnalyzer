package usecases

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/infrastructure"
)

type PredictUseCase interface {
	Predict(req domain.PredictRequest) (*domain.PredictResponse, error)
}

type predictUseCase struct {
	api *infrastructure.APIClient
}

func NewPredictUseCase(api *infrastructure.APIClient) PredictUseCase {
	return &predictUseCase{api: api}
}

func (u *predictUseCase) Predict(req domain.PredictRequest) (*domain.PredictResponse, error) {
	var res domain.PredictResponse
	err := u.api.PostJSON("/api/predict", req, &res)
	return &res, err
}
