package api

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendEmail 发送邮件
func SendEmail(ctx *gin.Context) {
	var userSend service.EmailService
	claims, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	err := ctx.ShouldBind(&userSend)
	if err == nil {
		res := userSend.Send(ctx.Request.Context(), claims.ID)
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// ValidEmail 验证邮箱
func ValidEmail(ctx *gin.Context) {
	var validEmail service.ValidEmailService
	err := ctx.ShouldBind(&validEmail)
	if err == nil {
		res := validEmail.ValidEmail(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(e.SUCCESS, res)
	} else {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
