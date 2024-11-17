package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"capstone-backend/dto"
	"capstone-backend/internal/controller"
	"capstone-backend/internal/middleware"
)

func Setup(listen bool, existingConn *gorm.DB) (*http.Server, *gin.Engine) {
	router := gin.New()
	router.Use(middleware.LoggerMiddleware())
	router.Use(gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, err any) {
		e, _ := err.(error)
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{Error: e.Error()})
	}))

	// Cors
	corsConfig := cors.DefaultConfig() // set CORS
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"User-Agent",
		"Referrer",
		"Host",
		"Token",
		"Authorization",
		"Cache-Control",
		"X-Request-ID",
		"Content-Disposition",
		"Pragma",
		"Expires",
		"Strict-Transport-Security",
		"X-Frame-Options",
		"X-XSS-Protection",
		"X-Content-Type-Options",
		"Content-Security-Policy",
	}
	corsConfig.AddAllowMethods("OPTIONS")
	router.Use(cors.New(corsConfig))

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "5000"
	}
	
	// init routes
	controller.InitRoutes(router, existingConn)

	// Serve app
	srv := &http.Server{
		Addr:         ":" + serverPort,
		Handler:      router,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	log.Printf("listening on %s\n", srv.Addr)

	if !listen {
		return srv, router
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return nil, nil
}
