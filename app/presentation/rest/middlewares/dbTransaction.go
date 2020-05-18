/**
 * Created by VoidArtanis on 11/2/2017
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/robsonmrsp/rest-security-go/app/data/repository"
)

// DBTransactionMidleware ...
func DBTransactionMidleware(dbHelper *repository.DataBaseHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbHelper.BeginTransaction()
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			dbHelper.RollbackTransaction()
			return
		}
		dbHelper.CommitTransaction()
	}
}
