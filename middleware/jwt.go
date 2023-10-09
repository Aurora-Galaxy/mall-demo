package middleware

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT 验证token
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == " " {
			code = e.TokenIsNULL
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
