package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wuzehv/ginframework/middleware"
	"github.com/wuzehv/ginframework/util/config"
	"log"
	"path/filepath"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.App.RunMode)

	router := gin.New()

	router.Use(middleware.Log())

	router.Use(gin.Recovery())

	router.LoadHTMLFiles(loadTemplates("template")...)

	construct(router)

	return router
}

func loadTemplates(templatesDir string) []string {
	other, err := filepath.Glob(templatesDir + "/**/*.html")
	if err != nil {
		log.Fatalf("load template error: %v\n", err)
	}

	admin, err := filepath.Glob(templatesDir + "/**/**/*.html")
	if err != nil {
		log.Fatalf("load template error: %v\n", err)
	}

	return append(admin, other...)
}
