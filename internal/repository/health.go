package repository

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type healthRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

type HealthRepository interface {
	Tx(db *gorm.DB) HealthRepository

	GetHealth(ctx context.Context) error
}

func NewHealthRepository(logger *zap.Logger, db *gorm.DB) healthRepository {
	return healthRepository{logger, db}
}

func (r healthRepository) Tx(db *gorm.DB) HealthRepository {
	if db == nil {
		return r
	}
	r.db = db
	return r
}

func (r healthRepository) GetHealth(ctx context.Context) error {
	return nil
}
