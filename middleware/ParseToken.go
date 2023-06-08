package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"personal_blog/db"
	"personal_blog/utility"
	"strings"
)

// ParseToken TODO 解析token中间件
func ParseToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token from another example.  This token is expired
		tokens := c.GetHeader("Authorization")
		if tokens == "" {
			c.Abort()
			panic("token不能为空！")
		}
		tokenString := strings.Split(tokens, "Bearer ")
		fmt.Println(tokenString[len(tokenString)-1])
		ctx := context.Background()
		user := utility.MyCustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString[len(tokenString)-1], &user, func(token *jwt.Token) (interface{}, error) {
			return utility.MySigningKey, nil
		})
		if err != nil {
			c.Abort()
			panic("令牌格式不对！")
		}

		if token.Valid {
			fmt.Println("验证通过")
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			fmt.Println("That's not even a token")
			c.Abort()
			panic("这不是一个token")
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			// Invalid signature
			fmt.Println("Invalid signature")
			c.Abort()
			panic("无效的签名")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			c.Abort()
			panic("token失效或过期！")
		} else {
			fmt.Println("Couldn't handle this token:", err)
			c.Abort()
			panic("未知错误！")
		}

		c.Set("Identification", user.Identification)
		fmt.Println("username", user.Identification)
		result, err := db.Rdb.Get(ctx, user.Identification).Result()
		if err != nil {
			c.Abort()
			panic("token失效或过期！")
		}
		fmt.Println(result)
		if result != tokenString[len(tokenString)-1] {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "账号在其他地方登录，只允许一台设备登录！",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
