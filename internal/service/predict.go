package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"capstone-backend/dto"

	"go.uber.org/zap"
)

type PredictService interface {
	Predict(ctx context.Context, req dto.PredictRequest) (dto.PredictResponse, error)
	PredictManual(ctx context.Context, req dto.PredictManualRequest) (dto.PredictResponse, error)
}

type predictService struct {
	logger *zap.Logger
}

func NewPredictService(logger *zap.Logger) PredictService {
	return &predictService{logger}
}

func (s *predictService) callPredictEndpoint(url string, data interface{}) (dto.PredictResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		s.logger.Error("JSON marshal error", zap.Error(err))
		return dto.PredictResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		s.logger.Error("HTTP POST error", zap.Error(err))
		return dto.PredictResponse{}, err
	}
	defer resp.Body.Close()

	var predictResponse dto.PredictResponse
	if err := json.NewDecoder(resp.Body).Decode(&predictResponse); err != nil {
		s.logger.Error("Response decode error", zap.Error(err))
		return dto.PredictResponse{}, err
	}

	return predictResponse, nil
}

func (s *predictService) Predict(ctx context.Context, req dto.PredictRequest) (dto.PredictResponse, error) {
	// adjust the url later
	url := "http://localhost:8000/predict"
	return s.callPredictEndpoint(url, req)
}

func (s *predictService) PredictManual(ctx context.Context, req dto.PredictManualRequest) (dto.PredictResponse, error) {
	// adjust the url later
	url := "http://localhost:8000/predict_manual"
	return s.callPredictEndpoint(url, req)
}
