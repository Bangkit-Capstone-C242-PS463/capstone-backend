package service

import (
	"capstone-backend/internal/constants"
	"capstone-backend/internal/model"
	"capstone-backend/internal/repository"
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userService struct {
	logger *zap.Logger
	ur     repository.UserRepository
}

type UserService interface {
	Tx(db *gorm.DB) userService

	GetUserHistory(ctx context.Context) ([]model.History, error)
	CreateHistory(ctx context.Context, result string) error
	DeleteHistoryByID(ctx context.Context, ID string) error
}

func NewUserService(logger *zap.Logger, ur repository.UserRepository) userService {
	return userService{logger, ur}
}

func (s userService) Tx(db *gorm.DB) userService {
	if db == nil {
		s.logger.Error("transaction database not found")
		return s
	}
	s.ur = s.ur.Tx(db)
	return s
}

func (s userService) GetUserHistory(ctx context.Context) ([]model.History, error) {
	userID := ctx.Value(constants.CONTEXT_USERID_KEY).(string)
	return s.ur.GetUserHistoryByUserID(ctx, userID)
}

func (s userService) CreateHistory(ctx context.Context, result string) error {
	userID := ctx.Value(constants.CONTEXT_USERID_KEY).(string)

	history := model.History{
		ID: gonanoid.Must(),
		UserID: userID,
		Result: result,
	}

	return s.ur.CreateHistory(ctx, &history)
}

func (s userService) DeleteHistoryByID(ctx context.Context, ID string) error {
	return s.ur.DeleteHistoryByID(ctx, ID)
}