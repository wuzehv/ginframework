package main

import (
	"github.com/wuzehv/ginframework/router"
	"github.com/wuzehv/ginframework/util/config"
	_ "github.com/wuzehv/ginframework/util/journal"
	"log"
)

// @Title ginframework项目
// @Version 1.0
// @Description ginframework项目
// @Contact.name wuzehui
// @Contact.email
// @Host liangjun.work
// @BasePath /
func main() {
	r := router.InitRouter()
	if err := r.Run(config.App.Port); err != nil {
		log.Fatalf("server run error: %v\n", err)
	}
}
