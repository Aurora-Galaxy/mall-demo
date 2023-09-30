package routes

import (
	"gin_mall/api"
	"gin_mall/middleware"
	"gin_mall/pkg/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//创建了一个用于存储会话数据的Cookie存储对象,[]byte中的字段为用于签署会话Cookie的密钥
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(e.SUCCESS, "success")
		})
		v1.POST("/user/register", api.UserRegister)
	}
	return r
}
