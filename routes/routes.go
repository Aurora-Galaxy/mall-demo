package routes

import (
	"gin_mall/api"
	"gin_mall/middleware"
	"gin_mall/pkg/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//创建了一个用于存储会话数据的Cookie存储对象,[]byte中的字段为用于签署会话Cookie的密钥
	store := cookie.NewStore([]byte("something-very-secret"))
	//实现跨域操作
	r.Use(middleware.Cors())
	//创建一个session中间件，mySession为该中间件的名称，store用于将 session 数据存储在客户端的 cookie 中
	r.Use(sessions.Sessions("mySession", store))
	//用于将指定目录下的文件提供为静态文件服务
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(e.SUCCESS, "success")
		})
		v1.POST("/user/register", api.UserRegister) //注册
		v1.POST("/user/login", api.UserLogin)       //登录

		//商品操作
		v1.GET("/carousel", api.ListCarousel)         //商品轮播图
		v1.GET("/product/list", api.ListProduct)      //展示商品列表
		v1.POST("/product/search", api.SearchProduct) //展示商品列表
		v1.GET("/product/:id", api.GetInfoProduct)    //获取商品详细信息
		v1.GET("/productimg/:id", api.GetImgProduct)  //获取商品图片路径
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.POST("user/update", api.UserUpdates)    //更新信息
			authed.POST("user/sendmail", api.SendEmail)    //发送邮件
			authed.POST("user/valid-mail", api.ValidEmail) //验证邮箱
			authed.GET("user/money", api.GetMoney)         //显示余额

			//商品操作
			authed.POST("product/create", api.CreateProduct) //创建商品

		}
	}
	return r
}
