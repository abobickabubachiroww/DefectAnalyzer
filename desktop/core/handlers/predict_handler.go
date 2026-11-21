package handlers

import (
	"DefectAnalyzer/desktop/core/domain"
	"DefectAnalyzer/desktop/core/usecases"
)

type PredictHandler struct {
	usecase usecases.PredictUseCase
}

func NewPredictHandler(uc usecases.PredictUseCase) *PredictHandler {
	return &PredictHandler{usecase: uc}
}

func (h *PredictHandler) Predict(req domain.PredictRequest) (*domain.PredictResponse, error) {
	return h.usecase.Predict(req)
}
