package controller

import (
	"gin-demo/config"
	m "gin-demo/middleware"
	"gin-demo/model/bo"
	"gin-demo/model/req"
	"gin-demo/model/res"
	"gin-demo/pkg/e"
	jwtUtil "gin-demo/pkg/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

func InitUser(engine *gin.Engine) {
	r := engine.Group("/api/user")
	r.POST("/")
	r.POST("/login", login)
	r.PUT("/:id").Use(m.JWTAuth())
	r.GET("/:id").Use(m.JWTAuth())
}

func login(ctx *gin.Context) {
	loginReq := req.LoginReq{}
	err := ctx.BindJSON(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Error(e.ERROR, "資料錯誤"))
		return
	}

	userBo := bo.UserBo{}
	err = copier.Copy(&userBo, &loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Error(e.ERROR, "資料錯誤"))
		return
	}
	var token string
	token, err = jwtUtil.GenToken(userBo, config.Setting.WebConfig.JwtSecret)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Error(e.ERROR, "token 產生錯誤"))
		return
	}

	ctx.JSON(http.StatusOK, res.Success(token))
}
