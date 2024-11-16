package auth

import (
	"net/http"
	"errors"

	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"capstone-backend/internal/service"
	"capstone-backend/internal/constants"
	"capstone-backend/internal/middleware"
	"capstone-backend/dto"
	customErr "capstone-backend/internal/errors"
)

type Controller struct {
	logger *zap.Logger
	auth   service.AuthService
}

func NewController(
	logger *zap.Logger,
	auth service.AuthService,
) *Controller {
	return &Controller{logger, auth}
}

func (ctrl *Controller) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Error("failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "failed to bind json",
			Error:   err.Error(),
		})
		return
	}

	txHandle := c.MustGet(constants.CONTEXT_TRX_KEY).(*gorm.DB)
	if err := ctrl.auth.Tx(txHandle).SignUp(c, req); err != nil {
		if errors.Is(err, customErr.ErrUserExisted) {
			c.JSON(http.StatusConflict, dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to sign up",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ctrl *Controller) Login(c *gin.Context) {
	var login dto.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		ctrl.logger.Error("failed to bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "failed to bind json",
			Error:   err.Error(),
		})
		return
	}
	resp, err := ctrl.auth.Login(c, login)
	if err != nil {
		ctrl.logger.Error("failed to login", zap.Error(err))
		if errors.Is(err, customErr.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "failed to login",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) Setup(router *gin.RouterGroup, db *gorm.DB) {
	r := router.Group("/v1/auth")
	{
		r.POST("/signup", middleware.DBTransactionMiddleware(db), ctrl.SignUp)
		r.POST("/login", ctrl.Login)
	}
}