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
		user.GET("/articlelist", middleware.ParseToken(), controller.GetArticle)
		user.GET("/commentlist", middleware.ParseToken(), controller.GetCommentList)
		user.POST("/addcomments", middleware.ParseToken(), controller.AddComment)
	}

	api := r.Group("/api", middleware.ParseToken(), middleware.ParseApi())
	r.GET("/api/login", controller.ApiLogin)
	{
		api.POST("/addarticle", controller.AddArticle)
		api.GET("/updatearticle", controller.UpdateArticle)
		api.DELETE("/deletearticle", controller.DeleteArticle)

	}

	return r

}
