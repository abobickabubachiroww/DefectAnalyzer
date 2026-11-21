package usecases

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/infrastructure"
)

type EconomicsUseCase interface {
	Calculate(req domain.EconomicsRequest) (*domain.EconomicsResponse, error)
}

type economicsUseCase struct {
	api *infrastructure.APIClient
}

func NewEconomicsUseCase(api *infrastructure.APIClient) EconomicsUseCase {
	return &economicsUseCase{api: api}
}

func (u *economicsUseCase) Calculate(req domain.EconomicsRequest) (*domain.EconomicsResponse, error) {
	var res domain.EconomicsResponse
	// Делаем POST запрос к Python API
	err := u.api.PostJSON("/api/economics/calculate", req, &res)
	return &res, err
}
