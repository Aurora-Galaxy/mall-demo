package api

import (
	"gin_mall/pkg/e"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	err := ctx.ShouldBind(&userRegister)
	if err == nil {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	err := ctx.ShouldBind(&userLogin)
	if err == nil {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
