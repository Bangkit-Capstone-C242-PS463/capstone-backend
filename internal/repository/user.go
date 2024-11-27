package repository

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"capstone-backend/internal/model"
)

type userRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

type UserRepository interface {
	Tx(db *gorm.DB) UserRepository

	Create(ctx context.Context, user model.User) error
	GetOneByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserHistoryByUserID(ctx context.Context, userID string) ([]model.History, error)
	CreateHistory(ctx context.Context, history *model.History) error
	DeleteHistoryByID(ctx context.Context, id string) error
}

func NewUserRepository(logger *zap.Logger, db *gorm.DB) userRepository {
	return userRepository{logger, db}
}

func (r userRepository) Tx(db *gorm.DB) UserRepository {
	if db == nil {
		return r
	}
	r.db = db
	return r
}

func (r userRepository) Create(ctx context.Context, user model.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (u userRepository) GetOneByEmail(ctx context.Context, email string) (*model.User, error) {
	var result model.User
	query := u.db.WithContext(ctx)
	query = query.Where("LOWER(email) = LOWER(?)", email)

	if err := query.First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (u userRepository) GetUserHistoryByUserID(ctx context.Context, userID string) (histories []model.History, err error) {
	if err := u.db.WithContext(ctx).Where("user_id = ?", userID).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (u userRepository) CreateHistory(ctx context.Context, history *model.History) error {
	if err := u.db.WithContext(ctx).Create(history).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepository) DeleteHistoryByID(ctx context.Context, id string) error {
	if err := u.db.WithContext(ctx).Unscoped().Where("id = ?", id).Delete(&model.History{}).Error; err != nil {
		return err
	}
	
	return nil
}