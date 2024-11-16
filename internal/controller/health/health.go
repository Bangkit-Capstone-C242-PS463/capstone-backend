package health

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"capstone-backend/internal/service"
)

type Controller struct {
	logger *zap.Logger
	hs     service.HealthService
}

func NewController(
	logger *zap.Logger,
	hs service.HealthService,
) *Controller {
	return &Controller{logger, hs}
}

func (ctrl *Controller) GetServerHealth(c *gin.Context) {
	ctrl.logger.Info("Checking server health...")
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}

func (ctrl *Controller) Setup(router *gin.RouterGroup) {
	r := router.Group("/v1/health")
	{
		r.GET("", ctrl.GetServerHealth)
	}
}
