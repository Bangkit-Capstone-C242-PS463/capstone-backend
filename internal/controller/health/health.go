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

// GetServerHealth godoc
// @Summary      Get server health status
// @Description  Returns the current health status of the server.
// @Tags         health
// @Produce      json
// @Success      200  {object}  dto.ServerHealthResponse "Server is healthy"
// @Router       /v1/health [get]
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
