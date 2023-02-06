package main

import (
	"fmt"
	"gin-demo/config"
	"gin-demo/controller"
	"gin-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	config.Load()
	controller.InitRouter(server)
	repository.InitDB()
	setting := config.Setting
	server.Run(fmt.Sprintf(":%d", setting.WebConfig.Port))
}
