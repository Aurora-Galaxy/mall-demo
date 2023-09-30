package api

import (
	"gin_mall/pkg/e"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
