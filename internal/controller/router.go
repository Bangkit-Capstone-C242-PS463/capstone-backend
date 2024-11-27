package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"capstone-backend/internal/controller/health"
	"capstone-backend/internal/controller/auth"
	"capstone-backend/internal/controller/user"
	"capstone-backend/internal/db"
	"capstone-backend/internal/logger"
	"capstone-backend/internal/repository"
	"capstone-backend/internal/service"
)

func InitRoutes(router *gin.Engine, conn *gorm.DB) {
	log := logger.GetLogger()

	// Initialise database.
	db := db.New()
	if conn != nil {
		db = conn
	}

	r := router.Group("/api")

	// Initialise repositories.
	hr := repository.NewHealthRepository(log, db)
	ur := repository.NewUserRepository(log, db)

	// Initialise services.
	hs := service.NewHealthService(log, hr)
	as := service.NewAuthService(log, ur)
	us := service.NewUserService(log, ur)

	// Initialise controllers.
	hCtrl := health.NewController(log, hs)
	hCtrl.Setup(r)

	authCtrl := auth.NewController(log, as)
	authCtrl.Setup(r, db)

	userCtrl := user.NewController(log, us)
	userCtrl.Setup(r, db)
}
