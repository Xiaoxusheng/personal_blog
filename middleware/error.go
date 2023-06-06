package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ddd", recover())
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 1,
					"err":  err,
				})
			}
		}()
		c.Next()
	}
}
