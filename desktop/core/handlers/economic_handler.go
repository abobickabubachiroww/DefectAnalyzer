package handlers

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/usecases"
)

type EconomicsHandler struct {
	usecase usecases.EconomicsUseCase
}

func NewEconomicsHandler(uc usecases.EconomicsUseCase) *EconomicsHandler {
	return &EconomicsHandler{usecase: uc}
}

func (h *EconomicsHandler) CalculateEconomics(req domain.EconomicsRequest) (*domain.EconomicsResponse, error) {
	return h.usecase.Calculate(req)
}
