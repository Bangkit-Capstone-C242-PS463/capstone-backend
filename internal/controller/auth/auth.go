package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"capstone-backend/dto"
	"capstone-backend/internal/constants"
	customErr "capstone-backend/internal/errors"
	"capstone-backend/internal/middleware"
	"capstone-backend/internal/service"
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

// SignUp godoc
// @Summary      Sign up a new user
// @Description  Registers a new user in the system.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body     dto.SignUpRequest   true  "User registration details"
// @Success      200   {object} dto.SuccessResponse "Sign-up success status"
// @Failure      400   {object} dto.ErrorResponse   "Bad Request - Invalid input"
// @Failure      409   {object} dto.ErrorResponse   "Conflict - User already exists"
// @Failure      500   {object} dto.ErrorResponse   "Internal Server Error - Sign-up failed"
// @Router       /v1/auth/signup [post]
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

// Login godoc
// @Summary      User login
// @Description  Authenticates a user and returns a data upon successful login.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login  body     dto.LoginRequest  	true  "User login credentials"
// @Success      200    {object} dto.LoginResponse  "Login success, returns user data"
// @Failure      400    {object} dto.ErrorResponse  "Bad Request - Invalid input"
// @Failure      401    {object} dto.ErrorResponse  "Unauthorized - Invalid credentials"
// @Failure      404    {object} dto.ErrorResponse  "Not Found - User not found"
// @Router       /v1/auth/login [post]
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