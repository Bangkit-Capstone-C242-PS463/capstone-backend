package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"capstone-backend/internal/logger"
)

type responseWriterWrapper struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rw *responseWriterWrapper) Write(data []byte) (int, error) {
	rw.body.Write(data)
	return rw.ResponseWriter.Write(data)
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetLogger()

		// Generate a unique request ID
		requestID := uuid.New().String()
		c.Set("requestID", requestID)
		log = log.With(zap.String("requestID", requestID))

		startTime := time.Now()

		// Log request details (method, URL, query params, headers)
		log.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.String()),
			zap.Any("query_params", c.Request.URL.Query()),
			zap.String("remote_addr", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Any("headers", c.Request.Header),
		)

		// Capture response body and status
		responseBody := &bytes.Buffer{}
		writer := &responseWriterWrapper{
			ResponseWriter: c.Writer,
			body:           responseBody,
			statusCode:     http.StatusOK,
		}
		c.Writer = writer

		// Process the request
		c.Next()

		// Log response details
		log.Info("Request completed",
			zap.Int("status_code", writer.statusCode),
			zap.Int64("duration_ms", time.Since(startTime).Milliseconds()),
			zap.String("response_body", responseBody.String()),
		)
	}
}
