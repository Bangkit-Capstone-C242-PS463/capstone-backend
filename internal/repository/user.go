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
