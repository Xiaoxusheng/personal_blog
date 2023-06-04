package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 1,
					"err":  err,
				})
			}
		}()
		c.Next()
	}
}
