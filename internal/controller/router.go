package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"capstone-backend/docs"
	"capstone-backend/internal/controller/auth"
	"capstone-backend/internal/controller/health"
	"capstone-backend/internal/controller/predict"
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

	// Programmatically set swagger info
	docs.SwaggerInfo.Title = "InSight API"
	docs.SwaggerInfo.Description = "Comprehensive documentation for the InSight API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialise repositories.
	hr := repository.NewHealthRepository(log, db)
	ur := repository.NewUserRepository(log, db)

	// Initialise services.
	hs := service.NewHealthService(log, hr)
	as := service.NewAuthService(log, ur)
	us := service.NewUserService(log, ur)
	ps := service.NewPredictService(log)

	// Initialise controllers.
	hCtrl := health.NewController(log, hs)
	hCtrl.Setup(r)

	authCtrl := auth.NewController(log, as)
	authCtrl.Setup(r, db)

	userCtrl := user.NewController(log, us)
	userCtrl.Setup(r, db)

	pCtrl := predict.NewPredictController(log, ps)
	pCtrl.Setup(r)
}
