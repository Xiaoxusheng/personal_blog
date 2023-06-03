package router

import (
	"github.com/gin-gonic/gin"
	"personal_blog/controller"
	"personal_blog/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Error())
	user := r.Group("/user")

	{
		user.POST("/login", controller.Login)
		user.POST("/register", controller.Register)
	}

	return r

}
