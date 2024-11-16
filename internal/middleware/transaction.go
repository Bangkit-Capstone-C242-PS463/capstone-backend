package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"capstone-backend/internal/constants"
)

// StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware : to setup the database transaction middleware
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		defer func() {
			// Panic occurred, rollback transaction to prevent idle in transaction
			if r := recover(); r != nil {
				txHandle.Rollback()
				log.Printf("Panic occurred: %v", r)
			} else if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated, http.StatusAccepted}) {
				// Valid status code, commit transaction
				if err := txHandle.Commit().Error; err != nil {
					log.Print("trx commit error: ", err)
				}
			} else {
				// Invalid status code, rollback transaction
				txHandle.Rollback()
			}
		}()

		c.Set(constants.CONTEXT_TRX_KEY, txHandle)
		c.Next()
	}
}
