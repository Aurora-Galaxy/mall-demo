package api

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetMoney 获取用户钱数
func GetMoney(ctx *gin.Context) {
	var moneyService service.ShowMoneyService
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	err := ctx.ShouldBind(&moneyService)
	if err == nil {
		res := moneyService.ShowMoney(ctx.Request.Context(), claims.ID)
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
