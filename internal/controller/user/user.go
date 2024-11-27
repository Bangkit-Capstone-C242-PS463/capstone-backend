package user

import (
	"capstone-backend/dto"
	"capstone-backend/internal/middleware"
	"capstone-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Controller struct {
	logger *zap.Logger
	user   service.UserService
}

func NewController(
	logger *zap.Logger,
	user service.UserService,
) *Controller {
	return &Controller{logger, user}
}

func (ctrl *Controller) GetAllHistory(c *gin.Context) {
	resp, err := ctrl.user.GetUserHistory(c)
	if err != nil {
		ctrl.logger.Error("failed to get history", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to get history",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"history": resp,
	})
}

func (ctrl *Controller) DeleteHistory(c *gin.Context) {
	historyID := c.Query("history_id")
	if historyID == "" {
		ctrl.logger.Error("history_id is missing")
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "history_id is missing",
		})
		return
	}

	if err := ctrl.user.DeleteHistoryByID(c, historyID); err != nil {
		ctrl.logger.Error("failed to delete history", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "failed to delete history",
			Error:   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ctrl *Controller) Setup(router *gin.RouterGroup, db *gorm.DB) {
	r := router.Group("/v1/user", middleware.JWTMiddleware())
	{
		r.GET("/history", ctrl.GetAllHistory)
		r.DELETE("/history", ctrl.DeleteHistory)
	}
}
