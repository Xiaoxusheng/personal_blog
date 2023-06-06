package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"personal_blog/controller"
	docs "personal_blog/docs"
	"personal_blog/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()

	//捕获错误中间件
	r.Use(middleware.Error())

	r.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := r.Group("/user")

	{
		user.POST("/login", controller.Login)
		user.POST("/register", controller.Register)
		user.POST("/addtitle", middleware.ParseToken(), controller.AddTitle)
	}

	return r

}
