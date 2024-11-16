package service

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"capstone-backend/internal/repository"
)

type healthService struct {
	logger *zap.Logger
	hr     repository.HealthRepository
}

type HealthService interface {
	Tx(db *gorm.DB) healthService

	CheckServerHealth(ctx context.Context) error
}

func NewHealthService(logger *zap.Logger, hr repository.HealthRepository) healthService {
	return healthService{logger, hr}
}

func (s healthService) Tx(db *gorm.DB) healthService {
	if db == nil {
		return s
	}
	s.hr = s.hr.Tx(db)
	return s
}

func (s healthService) CheckServerHealth(ctx context.Context) error {
	return s.hr.GetHealth(ctx)
}
