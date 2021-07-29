package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/ginframework/app"
	"github.com/wuzehv/ginframework/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wuzehv/ginframework/doc"
)

func construct(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := router.Group("/")
	r.Use(middleware.Base(), middleware.Cors())
	{
		r.GET("/", app.Index)
	}
}
