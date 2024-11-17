package service

import (
	"context"
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"capstone-backend/dto"
	"capstone-backend/internal/constants"
	customErr "capstone-backend/internal/errors"
	"capstone-backend/internal/model"
	"capstone-backend/internal/repository"
	"capstone-backend/utils"
)

type authService struct {
	logger *zap.Logger
	ur     repository.UserRepository
}

type AuthService interface {
	Tx(db *gorm.DB) authService

	SignUp(ctx context.Context, req dto.SignUpRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

func NewAuthService(logger *zap.Logger, ur repository.UserRepository) authService {
	return authService{logger, ur}
}

func (s authService) Tx(db *gorm.DB) authService {
	if db == nil {
		s.logger.Error("transaction database not found")
		return s
	}
	s.ur = s.ur.Tx(db)
	return s
}

func (s authService) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	s.logger.Info("user signup", zap.String("username", req.Username))

	// Check if user exists
	existingUser, errUser := s.ur.GetOneByUsername(ctx, req.Username)
	if errUser != nil && !errors.Is(errUser, gorm.ErrRecordNotFound) {
		s.logger.Error("failed to check username", zap.Error(errUser))
		return errUser
	}
	if existingUser != nil {
		s.logger.Error("user existed", zap.Error(errUser))
		return customErr.ErrUserExisted
	}

	userID := gonanoid.Must()

	// Hash password
	pw, salt, err := utils.CreateHashedPassword(req.Password, constants.PASSWORD_SALT_LENGTH)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return err
	}

	newUser := model.User{
		ID:           userID,
		Username:     req.Username,
		Name:         req.Name,
		Password:     &pw,
		PasswordSalt: &salt,
	}

	if err = s.ur.Create(ctx, newUser); err != nil {
		s.logger.Error("failed to create new user", zap.Error(err))
		return err
	}

	return nil
}

func (s authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	s.logger.Info("user login", zap.String("username", req.Username))

	result, err := s.ur.GetOneByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("username not found", zap.Error(err))
			return nil, customErr.ErrUserNotFound
		}
		s.logger.Error("failed to get user", zap.Error(err))
		return nil, err
	}

	isPwValid := utils.VerifyPasswordWithSalt(*result.Password, req.Password, *result.PasswordSalt)
	if isPwValid {
		token, err := utils.GenerateToken(result.ID)
		if err != nil {
			s.logger.Error("failed to generate token", zap.Error(err))
			return nil, err
		}

		return &dto.LoginResponse{
			ID:          result.ID,
			Username:    result.Username,
			Name:        result.Name,
			AccessToken: token,
		}, nil

	} else {
		s.logger.Error("password does not match")
		return nil, fmt.Errorf("password does not match")
	}
}
