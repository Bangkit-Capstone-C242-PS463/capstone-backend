package predict

import (
	"net/http"

	"capstone-backend/dto"
	"capstone-backend/internal/middleware"
	"capstone-backend/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PredictController struct {
	logger         *zap.Logger
	predictService service.PredictService
}

func NewPredictController(logger *zap.Logger, predictService service.PredictService) *PredictController {
	return &PredictController{
		logger:         logger,
		predictService: predictService,
	}
}

// Predict godoc
// @Summary      Predict disease based on user story
// @Description  Predicts the disease based on the provided user story.
// @Tags         prediction
// @Accept       json
// @Produce      json
// @Param        user_story  body      dto.PredictRequest      true  "User story for prediction"
// @Success      200         {object}  dto.PredictResponse     "Prediction successful"
// @Failure      400         {object}  dto.ErrorResponse       "Bad Request - Invalid input"
// @Failure      500         {object}  dto.ErrorResponse       "Internal Server Error - Failed to predict"
// @Router       /v1/predict [post]
func (ctrl *PredictController) Predict(c *gin.Context) {
	var req dto.PredictRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Error("Invalid PredictRequest", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	response, err := ctrl.predictService.Predict(c, req)
	if err != nil {
		ctrl.logger.Error("Prediction service error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to predict",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// PredictManual godoc
// @Summary      Predict disease based on symptoms
// @Description  Predicts the disease based on the provided symptoms.
// @Tags         prediction
// @Accept       json
// @Produce      json
// @Param        symptoms  body      dto.PredictManualRequest  true  "Symptoms for prediction"
// @Success      200       {object}  dto.PredictResponse       "Prediction successful"
// @Failure      400       {object}  dto.ErrorResponse         "Bad Request - Invalid input"
// @Failure      500       {object}  dto.ErrorResponse         "Internal Server Error - Failed to predict"
// @Router       /v1/predict_manual [post]
func (ctrl *PredictController) PredictManual(c *gin.Context) {
	var req dto.PredictManualRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ctrl.logger.Error("Invalid PredictManualRequest", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	response, err := ctrl.predictService.PredictManual(c, req)
	if err != nil {
		ctrl.logger.Error("Prediction service error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to predict",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *PredictController) Setup(router *gin.RouterGroup) {
	r := router.Group("/v1/predict", middleware.JWTMiddleware())
	{
		r.POST("/", ctrl.Predict)
		r.POST("/manual", ctrl.PredictManual)
	}
}
