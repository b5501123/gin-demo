package controller

import "github.com/gin-gonic/gin"

func InitRouter(engine *gin.Engine) {
	engine.Use(gin.Logger()) // 日志
	engine.Use(gin.Recovery())
	InitUser(engine)
}
