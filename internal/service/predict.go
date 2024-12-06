// internal/service/predict_service.go

package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"capstone-backend/dto"
	"capstone-backend/internal/constants"

	"go.uber.org/zap"
)

type PredictService interface {
	Predict(ctx context.Context, req dto.PredictRequest) (dto.PredictResponse, error)
	PredictManual(ctx context.Context, req dto.PredictManualRequest) (dto.PredictResponse, error)
}

type predictService struct {
	logger      *zap.Logger
	userService UserService
}

func NewPredictService(logger *zap.Logger, userService UserService) PredictService {
	return &predictService{
		logger:      logger,
		userService: userService,
	}
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
	response, err := s.callPredictEndpoint(constants.MODEL_PREDICT_URL, req)
	if err != nil {
		return dto.PredictResponse{}, err
	}

	// Save history using UserService
	err = s.userService.CreateHistory(ctx, response.PredictedDisease)
	if err != nil {
		s.logger.Error("Failed to create history", zap.Error(err))
	}

	return response, nil
}

func (s *predictService) PredictManual(ctx context.Context, req dto.PredictManualRequest) (dto.PredictResponse, error) {
	response, err := s.callPredictEndpoint(constants.MODEL_PREDICT_MANUAL_URL, req)
	if err != nil {
		return dto.PredictResponse{}, err
	}

	// Save history using UserService
	err = s.userService.CreateHistory(ctx, response.PredictedDisease)
	if err != nil {
		s.logger.Error("Failed to create history", zap.Error(err))
	}

	return response, nil
}
