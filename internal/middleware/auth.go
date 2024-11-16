package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"capstone-backend/dto"
	"capstone-backend/internal/constants"
	"capstone-backend/utils"
)

// Verifies the JWT token and authorizes the user
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "authorization header not provided",
			})
			c.Abort()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "authorization header must be in Bearer format",
			})
			c.Abort()
			return
		}

		// Verify the token
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "invalid or expired token",
			})
			c.Abort()
			return
		}

		// Extract user_id
		if userID, ok := claims[constants.CONTEXT_USERID_KEY].(float64); ok {
			c.Set(constants.CONTEXT_USERID_KEY, int64(userID))
		} else {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "invalid token claims",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}