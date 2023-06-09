package middleware

import (
	"github.com/gin-gonic/gin"
	"personal_blog/models"
)

// 管理员验证
func ParseApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		f, ok := c.Get("Identification")
		if !ok {
			c.Abort()
			panic("用户不存在！")
		}
		s := models.GetByUsePad(f.(string))
		if !s {
			c.Abort()
			panic("非法操作！")
		}
		c.Next()

	}
}
