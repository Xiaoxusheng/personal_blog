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

	r.Use(middleware.RequetLimite())

	r.Use(middleware.Cors())

	r.Use(middleware.LimitIp())

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := r.Group("/user")

	{
		user.POST("/login", controller.Login)
		user.GET("/logout", middleware.ParseToken(), controller.Logout)
		user.POST("/register", controller.Register)
		user.POST("/addarticle", middleware.ParseToken(), controller.AddArticle)
		user.GET("/articlelist", middleware.ParseToken(), controller.GetArticle)
		user.DELETE("/deletearticle", middleware.ParseToken(), controller.DeleteArticle)
		user.GET("/updatearticle", middleware.ParseToken(), controller.UpdateArticle)
	}

	return r

}
